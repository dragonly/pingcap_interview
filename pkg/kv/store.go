package kv

type Record struct {
	Key   int    // 主键，排序字段
	Data  []byte // 数据
}

type RecordKeyHeap []Record

func (h RecordKeyHeap) Len() int           { return len(h) }
func (h RecordKeyHeap) Less(i, j int) bool { return h[i].Key < h[j].Key }
func (h RecordKeyHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }


func (h *RecordKeyHeap) Push(x interface{}) {
	*h = append(*h, x.(Record))
}
func (h *RecordKeyHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

type Store struct {
	Records []Record
}
