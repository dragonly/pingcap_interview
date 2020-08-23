package multi_core

import (
	"github.com/dragonly/pingcap_interview/pkg/kv"
	"github.com/dragonly/pingcap_interview/pkg/local"
	"sync"
)

func GetTopNByMaxHeap(records []kv.Record, topN int) []kv.Record {
	if len(records) < topN {
		return records
	}
	chunks := local.split(records)
	wg := sync.WaitGroup{}
	// 分段在每个 chunk 内原地做 topN，也就是最后留下来的 chunk 中，前 min(topN, len(chunk)) 为改 chunk 中的 TopN
	for _, chunk := range chunks {
		wg.Add(1)
		chunk := chunk
		go func() {
			defer wg.Done()
			local.GetTopNByMaxHeap(chunk, topN)
		}()
	}
	wg.Wait()
	// reduce 操作，从每个 chunk 中获取 TopN，计算出总的 TopN
	reducedTopN := make([]kv.Record, len(chunks)*topN)
	dstStep := local.min(topN, len(chunks[0]))
	for i, chunk := range chunks {
		copy(reducedTopN[i*dstStep:], chunk[:local.min(topN, len(chunk))])
	}
	return local.GetTopNByMaxHeap(reducedTopN, topN)
}
