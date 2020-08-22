package single_core

import (
	"github.com/dragonly/pingcap_interview/pkg/kv"
	"sort"
)

// GetTopNBaseline 作为 baseline，先排序再取 topN，用来检验其他内存版本算法的正确性
func GetTopNBaseline(records []kv.Record, topN int) []kv.Record {
	if len(records) < topN {
		return records
	}
	sort.Sort(kv.RecordKeyHeap(records))
	return records[:topN]
}
