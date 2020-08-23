package local

import (
	"github.com/dragonly/pingcap_interview/pkg/kv"
	"github.com/rs/zerolog/log"
	"runtime"
	"sync"
)

func split(records []kv.Record) [][]kv.Record {
	// TODO: 优化数据量小的情况，或许可以做个自适应
	cores := runtime.NumCPU()
	log.Info().Msgf("cpu cores: %d", cores)
	chunkSize := len(records) / cores
	ret := make([][]kv.Record, cores)
	for i := 0; i < cores-1; i++ {
		ret[i] = records[i*chunkSize : (i+1)*chunkSize]
	}
	ret[cores-1] = records[(cores-1)*chunkSize:]
	return ret
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

// GetTopNParallel 实现了并行版本的调度，传入单核版本算法作为参数，分段计算再合并
func GetTopNParallel(records []kv.Record, topN int, topNFn TopNSolver) []kv.Record {
	if len(records) < topN {
		return records
	}
	chunks := split(records)
	wg := sync.WaitGroup{}
	reducedLen := 0
	for _, chunk := range chunks {
		wg.Add(1)
		chunk := chunk
		reducedLen += min(topN, len(chunk))
		go func() {
			defer wg.Done()
			topNFn(chunk, topN)
			//chunk[0].Key = -1
		}()
	}
	wg.Wait()
	// reduce 操作，从每个 chunk 中获取 TopN，计算出总的 TopN
	reducedTopN := make([]kv.Record, reducedLen)
	dstStep := min(topN, len(chunks[0]))
	for i, chunk := range chunks {
		copy(reducedTopN[i*dstStep:], chunk[:min(topN, len(chunk))])
	}
	return topNFn(reducedTopN, topN)
}
