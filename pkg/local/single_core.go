package local

import (
	"container/heap"
	"github.com/dragonly/pingcap_interview/pkg/kv"
	"github.com/rs/zerolog/log"
	"sort"
)

// GetTopNBaseline 作为 baseline，先排序再取 topN，用来检验其他内存版本算法的正确性
func GetTopNBaseline(records []kv.Record, topN int) []kv.Record {
	if len(records) < topN {
		return records
	}
	sort.Sort(kv.SortByRecordKey(records))
	return records[:topN]
}

// GetTopNMaxHeap 在 records 的前 min(TopN, len(records)) 范围内原地建堆，因此会导致传入数据发生变化
func GetTopNMaxHeap(records []kv.Record, topN int) []kv.Record {
	if len(records) < topN {
		return records
	}
	h := kv.RecordKeyMaxHeap(records[:topN])
	heap.Init(&h)
	log.Debug().Msgf("init: %v", h)
	for _, r := range records[topN:] {
		if r.Key < h[0].Key {
			h[0].Assign(r)
			log.Debug().Msgf("replace: %v", h)
			heap.Fix(&h, 0)
			log.Debug().Msgf("fix: %v", h)
			//log.Debug().Msgf("push: %v %v", h, r)
			//heap.Push(&h, r)
			//x := heap.Pop(&h)
			//log.Debug().Msgf("pop:  %v %v", h, x)
		}
	}
	return h
}

func GetTopNQuickSelect(records []kv.Record, topN int) []kv.Record {
	if len(records) < topN {
		return records
	}
	QuickSelect(kv.SortByRecordKey(records), topN)
	return records[:topN]
}
