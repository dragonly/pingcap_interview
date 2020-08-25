package cluster

import (
	"context"
	"github.com/dragonly/pingcap_interview/pkg/kv"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"net"
)

type server struct {
	UnimplementedTopNServer
}

func (s *server) TopNInBlock(ctx context.Context, request *TopNInBlockRequest) (*TopNInBlockResponse, error) {
	log.Info().Interface("request", request).Msg("received request")
	records := kv.ReadRecordsFile(request.DataBlock.Filename, request.DataBlock.BlockIndex)

	pRecords := make([]*Record, len(records))
	for i, r := range records {
		pRecords[i] = new(Record)
		pRecords[i].Key = r.Key
		pRecords[i].Data = make([]byte, len(r.Data))
		copy(pRecords[i].Data, r.Data)
	}
	return &TopNInBlockResponse{
		Records: pRecords[:1],
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
