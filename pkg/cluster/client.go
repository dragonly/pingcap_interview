package cluster

import (
	"context"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"time"
)

func GetTopNKeysInRange(minKey, maxKey int64) {
	ctxDial, cancelDial := context.WithTimeout(context.Background(), time.Second)
	defer cancelDial()
	address := viper.GetString("cluster.master.dial.address")
	log.Info().Msgf("dialing address %s", address)
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
			Filename: viper.GetString("cluster.data.file.path"),
			BlockIndex:  0,
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
