package local

import (
	"github.com/dragonly/pingcap_interview/pkg/storage"
	"sort"
	"testing"
)

func TestGetTopNBaseline(t *testing.T) {
	var records []storage.Record
	for i := 0; i < 20; i++ {
		r := storage.Record{
			Key:  int64(19 - i),
			Data: nil,
		}
		records = append(records, r)
	}

	r1 := GetTopNBaseline(records, 10)
	if len(r1) != 10 {
		t.Errorf("len=%d", len(r1))
	}
	for i, r := range r1 {
		if r.Key != int64(i) {
			t.Errorf("wrong key, i=%d, key=%d, should be key=%d", i, r.Key, i)
		}
	}
}

func TestGetTopNQuickSelect(t *testing.T) {
	var records []storage.Record
	for i := 0; i < 20; i++ {
		r := storage.Record{
			Key:  int64(19 - i),
			Data: nil,
		}
		records = append(records, r)
	}

	r1 := GetTopNQuickSelect(records, 10)
	sort.Sort(storage.SortByRecordKey(r1))
	if len(r1) != 10 {
		t.Errorf("len=%d", len(r1))
	}
	for i, r := range r1 {
		if r.Key != int64(i) {
			t.Errorf("wrong key, i=%d, key=%d, should be key=%d", i, r.Key, i)
		}
	}
}

func TestGetTopNMaxHeap(t *testing.T) {
	var records []storage.Record
	for i := 0; i < 20; i++ {
		r := storage.Record{
			Key:  int64(19 - i),
			Data: nil,
		}
		records = append(records, r)
	}

	r1 := GetTopNMaxHeap(records, 10)
	sort.Sort(storage.SortByRecordKey(r1))
	if len(r1) != 10 {
		t.Errorf("len=%d", len(r1))
	}
	for i, r := range r1 {
		if r.Key != int64(i) {
			t.Errorf("wrong key, i=%d, key=%d, should be key=%d", i, r.Key, i)
		}
	}
}

// TestGetTopNMaxHeapWithKeyRange 测试 topN 大于和小于 key range 范围的情况
func TestGetTopNMaxHeapWithKeyRange(t *testing.T) {
	var records []storage.Record
	for i := 0; i < 20; i++ {
		r := storage.Record{
			Key:  int64(i),
			Data: nil,
		}
		records = append(records, r)
	}

	// 可选范围大于 topN
	r1 := GetTopNMaxHeapWithKeyRange(records, 10, 5, 9)
	sort.Sort(storage.SortByRecordKey(r1))
	if len(r1) != 5 {
		t.Errorf("len=%d", len(r1))
	}
	for i, r := range r1 {
		if r.Key != int64(5+i) {
			t.Errorf("wrong key, i=%d, key=%d, should be key=%d", i, r.Key, i+5)
		}
	}

	// 可选范围小于 topN
	r2 := GetTopNMaxHeapWithKeyRange(records, 10, 5, 20)
	sort.Sort(storage.SortByRecordKey(r2))
	if len(r2) != 10 {
		t.Errorf("len=%d", len(r2))
	}
	for i, r := range r2 {
		if r.Key != int64(5+i) {
			t.Errorf("wrong key, i=%d, key=%d, should be key=%d", i, r.Key, i+5)
		}
	}
}

// TestGetTopNMaxHeapWithKeyRangeRandom 测试随机 key 分布情况
func TestGetTopNMaxHeapWithKeyRangeRandom(t *testing.T) {
	records := storage.GenRecords(20, 20)
	r1 := GetTopNMaxHeapWithKeyRange(records, 10, 5, 20)
	sort.Sort(storage.SortByRecordKey(r1))
	//fmt.Println(r1)
	if len(r1) != 10 {
		t.Errorf("len=%d", len(r1))
	}
	for i, r := range r1 {
		if r.Key != int64(5+i) {
			t.Errorf("wrong key, i=%d, key=%d, should be key=%d", i, r.Key, i+5)
		}
	}
}
