package local

import (
	"fmt"
	"github.com/dragonly/pingcap_interview/pkg/kv"
	"math"
	"sort"
)

const (
	n    = 10000
	topN = 10
	maxKey = math.MaxInt64
	//maxKey = n * 2
)

func Run() {
	records := kv.GenRecords(n)
	records1 := make([]kv.Record, n)
	records2 := make([]kv.Record, n)
	records3 := make([]kv.Record, n)
	records4 := make([]kv.Record, n)
	records5 := make([]kv.Record, n)
	//records6 := make([]kv.Record, n)
	copyRecords(records1, records)
	copyRecords(records2, records)
	copyRecords(records3, records)
	copyRecords(records4, records)
	copyRecords(records5, records)
	//copyRecords(records6, records)

	//store := kv.Store{Records: records}
	result1 := GetTopNBaseline(records1, topN)
	result2 := GetTopNMaxHeap(records2, topN)
	result3 := GetTopNQuickSelect(records3, topN)
	result4 := GetTopNParallel(records4, topN, GetTopNBaseline)
	result5 := GetTopNParallel(records5, topN, GetTopNMaxHeap)
	//result6 := GetTopNParallel(records5, topN, GetTopNQuickSelect)
	sort.Sort(kv.SortByRecordKey(result1))
	sort.Sort(kv.SortByRecordKey(result2))
	sort.Sort(kv.SortByRecordKey(result3))
	sort.Sort(kv.SortByRecordKey(result4))
	sort.Sort(kv.SortByRecordKey(result5))
	//sort.Sort(kv.SortByRecordKey(result6))
	fmt.Printf("result1:\n%v\n", result1)
	fmt.Printf("result2:\n%v\n", result2)
	fmt.Printf("result3:\n%v\n", result3)
	fmt.Printf("result4:\n%v\n", result4)
	fmt.Printf("result5:\n%v\n", result5)
	//fmt.Printf("result6:\n%v\n", result6)
}
