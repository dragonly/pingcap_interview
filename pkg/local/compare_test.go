package local

import (
	"github.com/dragonly/pingcap_interview/pkg/kv"
	"github.com/dragonly/pingcap_interview/pkg/local/multi_core"
	"reflect"
	"testing"
)

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
	result2 := GetTopNByMaxHeap(records2, topN)
	result3 := multi_core.GetTopNBaseline(records3, topN)
	result4 := multi_core.GetTopNByMaxHeap(records4, topN)
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

func BenchmarkBaselineSingle(b *testing.B) {
	records := GenerateRandomRecords(n)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		GetTopNBaseline(records, topN)
	}
}

func BenchmarkMinHeapSingle(b *testing.B) {
	records := GenerateRandomRecords(10000000)
	topN := 10
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		GetTopNByMaxHeap(records, topN)
	}
}
