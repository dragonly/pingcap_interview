package local

import (
	"fmt"
	"github.com/dragonly/pingcap_interview/pkg/kv"
	"sort"
)

const (
	n    = 1000
	topN = 10
	//maxKey = math.MaxInt64
	maxKey = n * 2
)

func Run() {
	records := GenerateRandomRecords(n)
	records1 := make([]kv.Record, n)
	records2 := make([]kv.Record, n)
	records3 := make([]kv.Record, n)
	records4 := make([]kv.Record, n)
	copyRecords(records1, records)
	copyRecords(records2, records)
	copyRecords(records3, records)
	copyRecords(records4, records)
	//fmt.Printf("records1:\n%v\n", records1)
	//fmt.Printf("records2:\n%v\n", records2)
	//fmt.Printf("records3:\n%v\n", records3)
	//fmt.Printf("records4:\n%v\n", records4)

	//store := kv.Store{Records: records}
	result1 := GetTopNBaseline(records1, topN)
	result2 := GetTopNMaxHeap(records2, topN)
	result3 := GetTopNParallel(records3, topN, GetTopNBaseline)
	result4 := GetTopNParallel(records3, topN, GetTopNMaxHeap)
	sort.Sort(kv.SortByRecordKey(result1))
	sort.Sort(kv.SortByRecordKey(result2))
	sort.Sort(kv.SortByRecordKey(result3))
	sort.Sort(kv.SortByRecordKey(result4))
	fmt.Printf("result1:\n%v\n", result1)
	fmt.Printf("result2:\n%v\n", result2)
	fmt.Printf("result3:\n%v\n", result3)
	fmt.Printf("result4:\n%v\n", result4)
}
