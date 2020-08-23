package local

import (
	"github.com/dragonly/pingcap_interview/pkg/kv"
	"github.com/rs/zerolog"
	"reflect"
	"sort"
	"testing"
)

func init() {
	zerolog.SetGlobalLevel(zerolog.ErrorLevel)
}

func TestHeapResult(t *testing.T) {
	// 由于算法会原地修改数据，需要各自 copy 一份输入数据
	records := GenerateRandomRecords(n)
	records1 := make([]kv.Record, n)
	records2 := make([]kv.Record, n)
	records3 := make([]kv.Record, n)
	records4 := make([]kv.Record, n)
	copy(records1, records)
	copy(records2, records)
	copy(records3, records)
	copy(records4, records)
	result1 := GetTopNBaseline(records1, topN)
	result2 := GetTopNMaxHeap(records2, topN)
	result3 := GetTopNParallel(records3, topN, GetTopNBaseline)
	result4 := GetTopNParallel(records4, topN, GetTopNMaxHeap)
	sort.Sort(kv.SortByRecordKey(result1))
	sort.Sort(kv.SortByRecordKey(result2))
	sort.Sort(kv.SortByRecordKey(result3))
	sort.Sort(kv.SortByRecordKey(result4))
	if !reflect.DeepEqual(result1, result2) {
		t.Errorf("single core heap method returns wrong results")
	}
	if !reflect.DeepEqual(result1, result3) {
		t.Errorf("multi core sort method returns wrong results")
	}
	if !reflect.DeepEqual(result1, result4) {
		t.Errorf("multi core heap method returns wrong results")
	}
}

func BenchmarkLocal(b *testing.B) {
	records := GenerateRandomRecords(n)
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
}

//func BenchmarkMinHeapSingle(b *testing.B) {
//	records := GenerateRandomRecords(n)
//	topN := 10
//	b.ResetTimer()
//	for i := 0; i < b.N; i++ {
//		GetTopNMaxHeap(records, topN)
//	}
//}
//
//func BenchmarkBaselineMulti(b *testing.B) {
//	records := GenerateRandomRecords(n)
//	b.ResetTimer()
//	for i := 0; i < b.N; i++ {
//		GetTopNParallel(records, topN, GetTopNBaseline)
//	}
//}
//
//func BenchmarkMinHeapMulti(b *testing.B) {
//	records := GenerateRandomRecords(n)
//	topN := 10
//	b.ResetTimer()
//	for i := 0; i < b.N; i++ {
//		GetTopNParallel(records, topN, GetTopNMaxHeap)
//	}
//}
