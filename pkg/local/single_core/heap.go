package single_core

import (
	"container/heap"
	"github.com/dragonly/pingcap_interview/pkg/kv"
)

func GetTopNByMinHeap(records []kv.Record, topN int) []kv.Record {
	if len(records) < topN {
		return records
	}
	h := kv.RecordKeyHeap(records[:topN])
	heap.Init(&h)
	for _, r := range records[topN:] {
		h.Push(r)
		h.Pop()
	}
	return h
}
