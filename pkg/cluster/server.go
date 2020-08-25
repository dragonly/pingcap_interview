package cluster

import (
	"context"
	"github.com/dragonly/pingcap_interview/pkg/kv"
	"github.com/dragonly/pingcap_interview/pkg/local"
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

func (s *server) TopNInBlock(ctx context.Context, request *TopNInBlockRequest) (*TopNInBlockResponse, error) {
	log.Info().Interface("request", request).Msg("received request")
	if request.TopN <= 0 {
		return nil, status.Errorf(codes.InvalidArgument, "invalid topN=%d", request.TopN)
	}
	t0 := time.Now()
	records := kv.ReadRecordsFile(request.DataBlock.Filename, request.DataBlock.BlockIndex)
	t1 := time.Now()
	topNRecords := local.GetTopNMaxHeap(records, int(request.TopN))
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

func StartServer() {
	address := viper.GetString("cluster.mapper.listen.address")
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
