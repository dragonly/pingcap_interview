package local

import (
	"fmt"
	"github.com/dragonly/pingcap_interview/pkg/kv"
	"github.com/rs/zerolog/log"
	"math/rand"
)

// GenerateRandomRecords 生成 Key 唯一且随机的包含 n 个 Record 的数组，data 数据随机
func GenerateRandomRecords(n int) []kv.Record {
	rand.Seed(0)
	dataLen := 1
	ret := make([]kv.Record, n)
	data := make([]byte, dataLen)
	existingKeys := make(map[int]struct{}, n)
	var record kv.Record
	for i := 0; i < n; i++ {
		var key int
		log.Debug().Msgf("%v", existingKeys)
		for {
			key = rand.Int() % maxKey
			if _, exist := existingKeys[key]; !exist {
				existingKeys[key] = struct{}{}
				break
			}
			log.Debug().Msgf("colliding key: %d", key)
		}
		if n, err := rand.Read(data); err != nil || n != dataLen {
			panic(fmt.Sprintf("err: %s, n: %d", err, n))
		}
		data1 := make([]byte, dataLen)
		copy(data1, data)
		record = kv.Record{Key: int64(key), Data: data1}
		ret[i] = record
	}
	return ret
}

func copyRecords(dst, src []kv.Record) {
	for i, _ := range src {
		dst[i].Key = src[i].Key
		copy(dst[i].Data, src[i].Data)
	}
}
