package local

import (
	"fmt"
	"github.com/dragonly/pingcap_interview/pkg/storage"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"math"
	"reflect"
	"sort"
	"testing"
)

func init() {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	log.Logger = log.With().Caller().Logger()
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

func TestHeapResult(t *testing.T) {
	n := 1000000
	topN := 10
	// 由于算法会原地修改数据，需要各自 copy 一份输入数据
	records := storage.GenRecords(n, math.MaxInt64)
	records1 := make([]storage.Record, n)
	records2 := make([]storage.Record, n)
	records3 := make([]storage.Record, n)
	records4 := make([]storage.Record, n)
	records5 := make([]storage.Record, n)
	records6 := make([]storage.Record, n)
	storage.CopyRecords(records1, records)
	storage.CopyRecords(records2, records)
	storage.CopyRecords(records3, records)
	storage.CopyRecords(records4, records)
	storage.CopyRecords(records5, records)
	storage.CopyRecords(records6, records)
	result1 := GetTopNBaseline(records1, topN)
	result2 := GetTopNMaxHeap(records2, topN)
	result3 := GetTopNQuickSelect(records3, topN)
	result4 := GetTopNParallel(records4, topN, GetTopNBaseline)
	result5 := GetTopNParallel(records5, topN, GetTopNMaxHeap)
	result6 := GetTopNParallel(records6, topN, GetTopNQuickSelect)
	sort.Sort(storage.SortByRecordKey(result1))
	sort.Sort(storage.SortByRecordKey(result2))
	sort.Sort(storage.SortByRecordKey(result3))
	sort.Sort(storage.SortByRecordKey(result4))
	sort.Sort(storage.SortByRecordKey(result5))
	sort.Sort(storage.SortByRecordKey(result6))
	if !reflect.DeepEqual(result1, result2) {
		t.Errorf("single core heap method returns wrong results")
	}
	if !reflect.DeepEqual(result1, result3) {
		t.Errorf("single core quick-select method returns wrong results")
	}
	if !reflect.DeepEqual(result1, result4) {
		t.Errorf("multi core sort method returns wrong results")
	}
	if !reflect.DeepEqual(result1, result5) {
		t.Errorf("multi core heap method returns wrong results")
	}
	if !reflect.DeepEqual(result1, result6) {
		t.Errorf("multi core quick-select method returns wrong results")
	}
}

func BenchmarkLocal(b *testing.B) {
	n := 1000000
	topN := 10
	records := storage.GenRecords(n, math.MaxInt64)
	b.ResetTimer()
	b.Run("BaselineSingle", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			GetTopNBaseline(records, topN)
		}
	})
	b.Run("MaxHeapSingle", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			GetTopNMaxHeap(records, topN)
		}
	})
	b.Run("QuickSelectSingle", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			GetTopNQuickSelect(records, topN)
		}
	})
	b.Run("BaselineMulti", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			GetTopNParallel(records, topN, GetTopNBaseline)
		}
	})
	b.Run("MaxHeapMulti", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			GetTopNParallel(records, topN, GetTopNMaxHeap)
		}
	})
	b.Run("QuickSelectMulti", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			GetTopNParallel(records, topN, GetTopNQuickSelect)
		}
	})
}
