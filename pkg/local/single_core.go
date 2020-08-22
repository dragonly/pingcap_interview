package local

import (
	"fmt"
	"github.com/dragonly/pingcap_interview/pkg/kv"
	"github.com/dragonly/pingcap_interview/pkg/local/single_core"
	"math/rand"
)

// generateRandomRecords 生成 Key 唯一且随机的包含 n 个 Record 的数组，data 数据随机
func generateRandomRecords(n int) []kv.Record {
	ret := make([]kv.Record, n)
	data := make([]byte, 10)
	existingKeys := make(map[int]struct{}, n)
	var record kv.Record
	for i := 0; i < n; i++ {
		var key int
		for {
			key = rand.Int()
			if _, exist := existingKeys[key]; !exist {
				break
			}
			fmt.Printf("colliding key: %d", key)
		}
		if n, err := rand.Read(data); err != nil || n != 10 {
			panic(fmt.Sprintf("err: %s, n: %d", err, n))
		}
		data1 := make([]byte, 10)
		copy(data1, data)
		record = kv.Record{Key: key, Data: data1}
		ret[i] = record
	}
	return ret
}

func run(s *kv.Store, topN int, getTopN TopNSolver) []kv.Record {
	return getTopN(s.Records, topN)
}

func Run() {
	records := generateRandomRecords(100)
	fmt.Println(len(records))
	//fmt.Println(records)
	store := kv.Store{Records: records}
	topN := 10
	result1 := run(&store, topN, single_core.GetTopNBaseline)
	result2 := run(&store, topN, single_core.GetTopNByMinHeap)
	fmt.Printf("result1:\n%v\n", result1)
	fmt.Printf("result2:\n%v\n", result2)
}
