package cluster

import (
	"context"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"time"
)

func GetTopNKeysInRange(minKey, maxKey, topN int64) {
	address := viper.GetString("cluster.master.dial.address")
	filename := viper.GetString("cluster.data.file.path")
	log.Info().
		Str("address", address).
		Str("data block filename", filename).
		Msg("run GetTopNKeysInRange")

	ctxDial, cancelDial := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancelDial()
	conn, err := grpc.DialContext(ctxDial, address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatal().Msgf("grpc.Dial: %v", err)
	}
	defer conn.Close()
	client := NewTopNClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	resp, err := client.TopNInBlock(ctx, &TopNInBlockRequest{
		DataBlock: &DataBlock{
			Filename:   filename,
			BlockIndex: 0,
		},
		KeyRange: &KeyRange{
			MaxKey: maxKey,
			MinKey: minKey,
		},
		TopN: topN,
	})
	if resp != nil {
		for _, r := range resp.Records {
			r.Data = r.Data[:10]
		}
	}
	log.Info().Interface("response", resp).Msg("receive from mapper")
	if err != nil {
		log.Fatal().Msgf("client.TopNInBlock error: %v", err)
	}
}
