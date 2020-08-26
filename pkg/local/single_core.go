package local

import (
	"container/heap"
	"github.com/dragonly/pingcap_interview/pkg/storage"
	"sort"
)

// GetTopNBaseline 作为 baseline，先排序再取 topN，用来检验其他内存版本算法的正确性
func GetTopNBaseline(records []storage.Record, topN int) []storage.Record {
	if len(records) < topN {
		return records
	}
	sort.Sort(storage.SortByRecordKey(records))
	return records[:topN]
}

// GetTopNMaxHeap 在 records 的前 min(TopN, len(records)) 范围内原地建堆，因此会导致传入数据发生变化
func GetTopNMaxHeap(records []storage.Record, topN int) []storage.Record {
	if len(records) < topN {
		return records
	}
	h := storage.RecordKeyMaxHeap(records[:topN])
	heap.Init(&h)
	//log.Debug().Msgf("init: %v", h)
	for _, r := range records[topN:] {
		if r.Key < h[0].Key {
			h[0].Assign(r)
			//log.Debug().Msgf("replace: %v", h)
			heap.Fix(&h, 0)
			//log.Debug().Msgf("fix: %v", h)
		}
	}
	return h
}

// GetTopNMaxHeapWithKeyRange 只返回 [minKey, maxKey] 范围内的 topN
func GetTopNMaxHeapWithKeyRange(records []storage.Record, topN int, minKey, maxKey int64) []storage.Record {
	var recordsTopN []storage.Record
	recordsInRange := 0
	var recordsRemainIndexStart int
	for i, r := range records {
		if r.Key >= minKey && r.Key <= maxKey {
			recordsTopN = append(recordsTopN, r)
			recordsInRange++
			if recordsInRange == topN {
				recordsRemainIndexStart = i + 1
				break
			}
		}
	}
	if recordsInRange < topN {
		return recordsTopN
	}
	h := storage.RecordKeyMaxHeap(recordsTopN)
	heap.Init(&h)
	for _, r := range records[recordsRemainIndexStart:] {
		if (r.Key >= minKey && r.Key <= maxKey) && r.Key < h[0].Key {
			h[0].Assign(r)
			heap.Fix(&h, 0)
		}
	}
	return h
}

func GetTopNQuickSelect(records []storage.Record, topN int) []storage.Record {
	if len(records) < topN {
		return records
	}
	QuickSelect(storage.SortByRecordKey(records), topN)
	return records[:topN]
}
