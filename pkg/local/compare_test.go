package local

import (
	"github.com/dragonly/pingcap_interview/pkg/kv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"reflect"
	"sort"
	"testing"
)

func init() {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	log.Logger = log.With().Caller().Logger()
}

func TestHeapResult(t *testing.T) {
	// 由于算法会原地修改数据，需要各自 copy 一份输入数据
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
	//if !reflect.DeepEqual(result1, result6) {
	//	t.Errorf("multi core quick-select method returns wrong results")
	//}
}

func BenchmarkLocal(b *testing.B) {
	records := kv.GenRecords(n)
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
