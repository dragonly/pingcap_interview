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
	"net"
	"time"
)

type server struct {
	UnimplementedTopNServer
}

// TopNInBlock 读取一个文件 block，计算其中的 topN，并通过 gRPC 返回结果
func (s *server) TopNInBlock(ctx context.Context, request *TopNInBlockRequest) (*TopNInBlockResponse, error) {
	log.Info().Interface("request", request).Msg("TopNInBlock received request")
	topN := request.TopN
	//minKey := request.KeyRange.MinKey
	//maxKey := request.KeyRange.MaxKey
	if topN <= 0 {
		return nil, status.Errorf(codes.InvalidArgument, "invalid topN=%d", topN)
	}
	// TODO: 可以遵守一下 deadline，避免浪费计算资源
	deadline, hasDeadline := ctx.Deadline()
	log.Debug().Msgf("time remains for request: %d", deadline.Sub(time.Now()))
	t0 := time.Now()
	// TODO: optimization: 先只加载 key，计算出结果后再读取响应的 topN 返回结果
	records := storage.ReadRecordsFile(request.DataBlock.Filename, request.DataBlock.BlockIndex)
	if hasDeadline && time.Now().Sub(deadline) > 0 {
		return nil, status.Errorf(codes.DeadlineExceeded, "deadline exceeded, skip calculation")
	}
	t1 := time.Now()
	topNRecords := local.GetTopNMaxHeap(records, int(topN))
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
		Int("records in block", len(records)).
		Int("topN records in block", len(topNRecords)).
		Dict("time_us", zerolog.Dict().
			Int64("kv.ReadRecordsFile()", t1.Sub(t0).Microseconds()).
			Int64("local.GetTopNMaxHeap()", t2.Sub(t1).Microseconds()).
			Int64("copy result", t3.Sub(t2).Microseconds())).
		Msg("return topN records in block")
	return &TopNInBlockResponse{
		Records: pRecords,
	}, nil
}

// TopNAll 一次性计算所有文件 block 的 topN，用于验证计算正确性
func (s *server) TopNAll(ctx context.Context, request *TopNInBlockRequest) (*TopNInBlockResponse, error) {
	log.Info().Interface("request", request).Msg("TopNAll received request")
	blockNum := viper.GetInt("cluster.data.file.blockNum")
	topN := request.TopN
	if topN <= 0 {
		return nil, status.Errorf(codes.InvalidArgument, "invalid topN=%d", topN)
	}

	var topNRecords []storage.Record
	for i := 0; i < blockNum; i++ {
		records := storage.ReadRecordsFile(request.DataBlock.Filename, request.DataBlock.BlockIndex)
		tmp := local.GetTopNMaxHeap(records, int(topN))
		for _, r := range tmp {
			topNRecords = append(topNRecords, r)
		}
	}
	topNRecords = local.GetTopNMaxHeap(topNRecords, int(topN))
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
	RegisterTopNServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatal().Msgf("failed to serve: %v", err)
	}
}
