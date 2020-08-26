package cluster

import (
	"context"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"time"
)

type Job struct {
	ID      int
	Request *TopNInBlockRequest
}

type Result struct {
	Job      Job
	Response *TopNInBlockResponse
	Error    error
}

type Worker struct {
	WorkerChan chan Worker // 将自己放入 WorkerChan，表示空闲
	JobChan    chan Job    // 从 JobChan 中接收任务并执行
	ResultChan chan Result // 将任务执行结果放入 ResultChan
	StopChan   chan bool   // 接收到 StopChan 的信号后，停止当前 Worker
	Client     *TopNClient // gRPC client，请求对应 mapper 节点执行实际计算任务
}

func DoRequest(client TopNClient, req *TopNInBlockRequest) (*TopNInBlockResponse, error) {
	timeoutMs := viper.GetInt("cluster.master.request.timeout")
	log.Debug().Msgf("start DoRequest() with timeout=%dms", timeoutMs)
	//fmt.Println(time.Duration(timeoutMs)*time.Millisecond)
	//fmt.Println(30000*time.Millisecond)
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeoutMs)*time.Millisecond)
	defer cancel()
	resp, err := client.TopNInBlock(ctx, req, grpc.MaxCallRecvMsgSize(1024*1024*1024))
	if resp != nil {
		for _, r := range resp.Records {
			r.Data = r.Data[:10]
			r.Data = nil
		}
	}
	if err != nil {
		log.Error().Msgf("client.TopNInBlock error: %v", err)
	}
	return resp, err
}

func (w *Worker) Start() {
	go func() {
		for {
			w.WorkerChan <- *w
			select {
			case job := <-w.JobChan:
				log.Debug().Int("id", job.ID).Int("channels", len(w.JobChan)).Msg("worker start job")
				resp, err := DoRequest(*w.Client, job.Request)
				if err != nil {
					log.Error().Int("id", job.ID).Str("error", err.Error()).Msg("worker job failed")
				} else {
					log.Debug().Int("id", job.ID).Msg("worker job success")
				}
				result := Result{
					Job:      job,
					Response: resp,
					Error:    err,
				}
				w.ResultChan <- result
				log.Debug().Int("id", job.ID).Int("channels", len(w.ResultChan)).Msg("worker sent result")
			case <-w.StopChan:
				log.Info().Msgf("worker exit: %v", w)
				return
			}
		}
	}()
}

func (w *Worker) Stop() {
	log.Info().Msgf("stop worker: %v", w)
	w.StopChan <- true
}

type Dispatcher struct {
	clients           []TopNClient
	WorkerChan        chan Worker // 从 WorkerChan 中获取空闲的 Worker，将任务分配给它
	JobChan           chan Job    // 从 driver 接收 Job，用于分配给 Worker
	JobRescheduleChan chan Result // 从 JobRescheduleChan 中检查 Job 结果，决定是否重新调度计算
	JobResultChan     chan Result // 向 driver 发送 Job 结果
	StopChan          chan bool   // 停止信号 channel
}

func NewDispatcher(clients []TopNClient, jobNum int) Dispatcher {
	return Dispatcher{
		clients:           clients,
		WorkerChan:        make(chan Worker, len(clients)),
		JobChan:           make(chan Job, jobNum),
		JobRescheduleChan: make(chan Result, jobNum),
		JobResultChan:     make(chan Result, jobNum),
		StopChan:          make(chan bool),
	}
}

func (d *Dispatcher) Start() {
	var workers []Worker
	// 启动 workers
	for _, c := range d.clients {
		c := c
		worker := Worker{
			WorkerChan: d.WorkerChan,
			JobChan:    make(chan Job),
			ResultChan: d.JobRescheduleChan,
			StopChan:   make(chan bool),
			Client:     &c,
		}
		worker.Start()
		workers = append(workers, worker)
	}

	go func() {
		for {
			select {
			case job := <-d.JobChan: // 从外部获取任务
				log.Debug().Int("id", job.ID).Int("channels", len(d.JobChan)).Msg("dispatcher received job")
				worker := <-d.WorkerChan // 获取空闲 worker
				log.Debug().Int("id", job.ID).Int("channels", len(d.WorkerChan)).Msg("dispatcher got idle worker")
				worker.JobChan <- job // 分配任务给 worker
				log.Debug().Int("id", job.ID).Int("channels", len(worker.JobChan)).Msg("dispatcher dispatched job")
			case result := <-d.JobRescheduleChan:
				if result.Error != nil { // 失败任务重试
					// TODO: 处理 tcp 连接断开的情况，应该等待 gRPC 重连，可以在 Job 中加一段等待时间
					log.Error().Int("id", result.Job.ID).Msg("dispatcher re-dispatch job")
					d.JobChan <- result.Job
				} else {
					log.Info().Int("id", result.Job.ID).Msg("dispatcher return successful job")
					d.JobResultChan <- result
				}
			case <-d.StopChan:
				log.Info().Msg("dispatcher stopping workers")
				for _, w := range workers {
					w.Stop()
				}
				return
			}
		}
	}()
}
