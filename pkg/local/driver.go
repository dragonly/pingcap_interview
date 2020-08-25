package local

import (
	"fmt"
	"github.com/dragonly/pingcap_interview/pkg/storage"
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
	records := storage.GenRecords(n)
	records1 := make([]storage.Record, n)
	records2 := make([]storage.Record, n)
	records3 := make([]storage.Record, n)
	records4 := make([]storage.Record, n)
	records5 := make([]storage.Record, n)
	//records6 := make([]kv.Record, n)
	storage.CopyRecords(records1, records)
	storage.CopyRecords(records2, records)
	storage.CopyRecords(records3, records)
	storage.CopyRecords(records4, records)
	storage.CopyRecords(records5, records)
	//kv.CopyRecords(records6, records)

	//store := kv.Store{Records: records}
	result1 := GetTopNBaseline(records1, topN)
	result2 := GetTopNMaxHeap(records2, topN)
	result3 := GetTopNQuickSelect(records3, topN)
	result4 := GetTopNParallel(records4, topN, GetTopNBaseline)
	result5 := GetTopNParallel(records5, topN, GetTopNMaxHeap)
	//result6 := GetTopNParallel(records5, topN, GetTopNQuickSelect)
	sort.Sort(storage.SortByRecordKey(result1))
	sort.Sort(storage.SortByRecordKey(result2))
	sort.Sort(storage.SortByRecordKey(result3))
	sort.Sort(storage.SortByRecordKey(result4))
	sort.Sort(storage.SortByRecordKey(result5))
	//sort.Sort(kv.SortByRecordKey(result6))
	fmt.Printf("result1:\n%v\n", result1)
	fmt.Printf("result2:\n%v\n", result2)
	fmt.Printf("result3:\n%v\n", result3)
	fmt.Printf("result4:\n%v\n", result4)
	fmt.Printf("result5:\n%v\n", result5)
	//fmt.Printf("result6:\n%v\n", result6)
}
