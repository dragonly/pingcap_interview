package cluster

import (
	"context"
	"github.com/dragonly/pingcap_interview/pkg/kv"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"net"
)

const (
	address = "localhost:2333"
)

type server struct {
	UnimplementedTopNServer
}

func (s *server) TopNInBlock(ctx context.Context, request *TopNInBlockRequest) (*TopNInBlockResponse, error) {
	log.Info().Interface("request", request).Msg("received request")
	kv.ReadRecordsFile(request.DataBlock.Filename, request.DataBlock.BlockIndex)
	return &TopNInBlockResponse{
		Records: []*Record{},
	}, nil
}

func StartServer() {
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal().Msgf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	RegisterTopNServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatal().Msgf("failed to serve: %v", err)
	}
}
