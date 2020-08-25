package local

import (
	"fmt"
	"github.com/dragonly/pingcap_interview/pkg/kv"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"sort"
)

func Run() {
	n := viper.GetInt("local.data.record.num")
	topN := viper.GetInt("local.data.record.topN")
	log.Info().
		//Interface("local.data.record", viper.GetStringMap("local.data.record")).
		Int("n", n).
		Int("topN", topN).
		Msg("local algorithm test")
	records := kv.GenRecords(n)
	records1 := make([]kv.Record, n)
	records2 := make([]kv.Record, n)
	records3 := make([]kv.Record, n)
	records4 := make([]kv.Record, n)
	records5 := make([]kv.Record, n)
	//records6 := make([]kv.Record, n)
	kv.CopyRecords(records1, records)
	kv.CopyRecords(records2, records)
	kv.CopyRecords(records3, records)
	kv.CopyRecords(records4, records)
	kv.CopyRecords(records5, records)
	//kv.CopyRecords(records6, records)

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
