package cluster

import (
	"context"
	"github.com/dragonly/pingcap_interview/pkg/local"
	"github.com/dragonly/pingcap_interview/pkg/storage"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"math/rand"
	"net"
	_ "net/http/pprof"
	"runtime"
	"sort"
	"time"
)

type server struct {
	UnimplementedTopNServer
}

// TopNInBlock 读取一个文件 block，计算其中的 topN，并通过 gRPC 返回结果
func (s *server) TopNInBlock(ctx context.Context, request *TopNInBlockRequest) (*TopNInBlockResponse, error) {
	log.Info().Interface("request", request).Msg("TopNInBlock received request")
	topN := request.TopN
	minKey := request.KeyRange.MinKey
	maxKey := request.KeyRange.MaxKey
	blockIndex := request.DataBlock.BlockIndex
	filename := request.DataBlock.Filename
	failRate := request.FailRate

	rand.Seed(time.Now().Unix())
	if rand.Float32() < failRate {
		log.Error().Msg("sleep and abort with random failure")
		time.Sleep(time.Second)
		return nil, status.Errorf(codes.Aborted, "random failure")
	}
	if topN <= 0 {
		return nil, status.Errorf(codes.InvalidArgument, "invalid topN=%d", topN)
	}

	deadline, hasDeadline := ctx.Deadline()
	log.Debug().Msgf("time remains for request: %d", deadline.Sub(time.Now()))
	t0 := time.Now()
	// TODO: optimization: 先只加载 key，计算出结果后再读取相应的 topN 返回结果
	records := storage.ReadRecordsFile(filename, blockIndex)
	if hasDeadline && time.Now().Sub(deadline) > 0 {
		return nil, status.Errorf(codes.DeadlineExceeded, "deadline exceeded, skip calculation")
	}
	t1 := time.Now()
	topNRecords := local.GetTopNMaxHeapWithKeyRange(records, int(topN), minKey, maxKey)
	t2 := time.Now()
	pRecords := make([]*Record, len(topNRecords))
	for i, r := range topNRecords {
		pRecords[i] = new(Record)
		pRecords[i].Key = r.Key
		pRecords[i].Data = make([]byte, len(r.Data))
		copy(pRecords[i].Data, r.Data)
	}
	t3 := time.Now()
	log.Info().
		Int64("blockIndex", blockIndex).
		Str("filename", filename).
		Int("records in block", len(records)).
		Int("topN records in block", len(topNRecords)).
		Dict("time_us", zerolog.Dict().
			Int64("kv.ReadRecordsFile()", t1.Sub(t0).Microseconds()).
			Int64("local.GetTopNMaxHeap()", t2.Sub(t1).Microseconds()).
			Int64("copy result", t3.Sub(t2).Microseconds())).
		Msg("topN records in block")
	for _, r := range records {
		r.Data = nil
	}
	records = nil
	topNRecords = nil
	runtime.GC()
	return &TopNInBlockResponse{
		Records: pRecords,
	}, nil
}

// TopNAll 一次性计算所有文件 block 的 topN，用于验证计算正确性
func (s *server) TopNAll(ctx context.Context, request *TopNInBlockRequest) (*TopNInBlockResponse, error) {
	log.Info().Interface("request", request).Msg("TopNAll received request")
	blockNum := viper.GetInt("cluster.data.file.blockNum")
	topN := request.TopN
	minKey := request.KeyRange.MinKey
	maxKey := request.KeyRange.MaxKey
	//blockIndex := request.DataBlock.BlockIndex
	filename := request.DataBlock.Filename

	if topN <= 0 {
		return nil, status.Errorf(codes.InvalidArgument, "invalid topN=%d", topN)
	}

	var topNRecords []storage.Record
	for i := 0; i < blockNum; i++ {
		t0 := time.Now()
		records := storage.ReadRecordsFile(filename, int64(i))
		t1 := time.Now()
		tmp := local.GetTopNMaxHeapWithKeyRange(records, int(topN), minKey, maxKey)
		t2 := time.Now()
		log.Info().
			Int("blockIndex", i).
			Str("filename", filename).
			Int("records in block", len(records)).
			Int("topN records in block", len(topNRecords)).
			Dict("time_us", zerolog.Dict().
				Int64("kv.ReadRecordsFile()", t1.Sub(t0).Microseconds()).
				Int64("local.GetTopNMaxHeap()", t2.Sub(t1).Microseconds())).
			Msg("topN records in block")
		for _, r := range tmp {
			topNRecords = append(topNRecords, r)
		}
	}
	topNRecords = local.GetTopNMaxHeap(topNRecords, int(topN))
	sort.Sort(storage.SortByRecordKey(topNRecords))
	//fmt.Println(topNRecords)

	pRecords := make([]*Record, len(topNRecords))
	for i, r := range topNRecords {
		pRecords[i] = new(Record)
		pRecords[i].Key = r.Key
		pRecords[i].Data = make([]byte, len(r.Data))
		copy(pRecords[i].Data, r.Data)
	}
	return &TopNInBlockResponse{
		Records: pRecords,
	}, nil
}

func StartServer(serverIndex int) {
	address := viper.GetStringSlice("cluster.mapper.listen.addresses")[serverIndex]
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal().Msgf("failed to listen: %v", err)
	}
	log.Info().Msgf("listening on %s", address)
	s := grpc.NewServer()
	//s := grpc.NewServer(
	//	grpc.KeepaliveParams(keepalive.ServerParameters{
	//		MaxConnectionIdle: 5 * time.Minute,
	//	}),
	//)
	RegisterTopNServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatal().Msgf("failed to serve: %v", err)
	}
}
