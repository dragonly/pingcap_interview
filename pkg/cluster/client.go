package cluster

import (
	"context"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"time"
)

func GetTopNKeysInRange(minKey, maxKey int64) {
	ctxDial, cancelDial := context.WithTimeout(context.Background(), time.Second)
	defer cancelDial()
	conn, err := grpc.DialContext(ctxDial, address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatal().Msgf("grpc.Dial: %v", err)
	}
	defer conn.Close()
	client := NewTopNClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	resp, err := client.TopNInBlock(ctx, &TopNInBlockRequest{
		DataBlock: &DataBlock{
			Filename: "test",
			BlockId:  0,
		},
		KeyRange: &KeyRange{
			MaxKey: maxKey,
			MinKey: minKey,
		},
	})
	log.Info().Interface("response", resp).Msg("receive from mapper")
	if err != nil {
		log.Fatal().Msgf("client.TopNInBlock error: %v", err)
	}
}
