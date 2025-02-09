package storage

type Record struct {
	Key  int64    // 主键，排序字段
	Data []byte // 数据
}

func CopyRecords(dst, src []Record) {
	for i, r := range src {
		dst[i].Assign(r)
	}
}

func (r *Record) Assign(r1 Record) {
	r.Key = r1.Key
	r.Data = make([]byte, len(r1.Data))
	copy(r.Data, r1.Data)
}

type SortByRecordKey []Record

func (h SortByRecordKey) Len() int           { return len(h) }
func (h SortByRecordKey) Less(i, j int) bool { return h[i].Key < h[j].Key }
func (h SortByRecordKey) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

type RecordKeyMaxHeap []Record

func (h RecordKeyMaxHeap) Len() int           { return len(h) }
func (h RecordKeyMaxHeap) Less(i, j int) bool { return h[i].Key > h[j].Key }
func (h RecordKeyMaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *RecordKeyMaxHeap) Push(x interface{}) {
	*h = append(*h, x.(Record))
}

func (h *RecordKeyMaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

type Store struct {
	Records []Record
}
