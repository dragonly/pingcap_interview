package cluster

import (
	"context"
	"github.com/dragonly/pingcap_interview/pkg/local"
	"github.com/dragonly/pingcap_interview/pkg/storage"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"time"
)

func GetTopNKeysInRange(minKey, maxKey int64, topN int) {
	// 获取计算任务所需参数
	addresses := viper.GetStringSlice("cluster.master.dial.addresses")
	filename := viper.GetString("cluster.data.file.path")
	blockNum := viper.GetInt("cluster.data.file.blockNum")
	log.Info().
		Strs("addresses", addresses).
		Str("data block filename", filename).
		Msg("run GetTopNKeysInRange")

	// 连接所有计算节点
	log.Info().Msg("connecting calculating nodes")
	var clients []TopNClient
	for _, addr := range addresses {
		ctxDial, cancelDial := context.WithTimeout(context.Background(), time.Second)
		//goland:noinspection GoDeferInLoop
		defer cancelDial()
		conn, err := grpc.DialContext(ctxDial, addr, grpc.WithInsecure(), grpc.WithBlock())
		if err != nil {
			log.Error().Msgf("grpc.Dial failed for address %s: %v", addr, err)
			continue
		}
		//goland:noinspection GoDeferInLoop
		defer conn.Close()
		clients = append(clients, NewTopNClient(conn))
		log.Info().Msgf("client inialized for %s", addr)
	}
	if len(clients) == 0 {
		log.Fatal().Msg("no mapper service available")
	}

	// 构建计算任务列表
	var jobs []Job
	for i := 0; i < blockNum; i++ {
		job := Job{
			ID: i,
			Request: &TopNInBlockRequest{
				DataBlock: &DataBlock{
					Filename:   filename,
					BlockIndex: int64(i),
				},
				KeyRange: &KeyRange{
					MaxKey: maxKey,
					MinKey: minKey,
				},
				TopN: int64(topN),
			},
		}
		jobs = append(jobs, job)
	}

	// 调度计算任务
	log.Info().Msg("driver dispatching jobs")
	dispatcher := NewDispatcher(clients, len(jobs))
	dispatcher.Start()
	for _, job := range jobs {
		log.Debug().Int("id", job.ID).Int("channels", len(dispatcher.JobChan)).Msg("driver dispatching job")
		dispatcher.JobChan <- job
		log.Debug().Int("id", job.ID).Int("channels", len(dispatcher.JobChan)).Msg("driver dispatched job")
	}

	// 获取分块任务 topN，合并最终结果
	log.Info().Msg("driver reduce topN")
	var records []storage.Record
	for i := 0; i < len(jobs); i++ {
		result := <-dispatcher.JobResultChan
		for _, pRecord := range result.Response.Records {
			record := storage.Record{}
			record.Key = pRecord.Key
			record.Data = make([]byte, len(pRecord.Data))
			copy(record.Data, pRecord.Data)
			records = append(records, record)
		}
	}
	topNRecords := local.GetTopNBaseline(records, topN)

	// debug 用，只看 key
	for _, r := range topNRecords {
		r.Data = nil
	}

	log.Info().Interface("keys", topNRecords).Msgf("got top n records")
}
