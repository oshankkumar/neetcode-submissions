
type IntCounter struct {
	Val   int
	Count int
}

type IntCounterHeap []IntCounter

func (h IntCounterHeap) Vals() []int {
	keys := make([]int, len(h))
	for i, v := range h {
		keys[i] = v.Val
	}
	return keys
}

// Len implements heap.Interface.
func (h *IntCounterHeap) Len() int {
	return len(*h)
}

// Less implements heap.Interface.
func (h *IntCounterHeap) Less(i int, j int) bool {
	return (*h)[i].Count < (*h)[j].Count
}

// Pop implements heap.Interface.
func (h *IntCounterHeap) Pop() any {
	v := (*h)[h.Len()-1]
	*h = append(IntCounterHeap(nil), (*h)[:h.Len()-1]...)
	return v
}

// Push implements heap.Interface.
func (h *IntCounterHeap) Push(x any) {
	*h = append(*h, x.(IntCounter))
}

// Swap implements heap.Interface.
func (h *IntCounterHeap) Swap(i int, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

var _ heap.Interface = (*IntCounterHeap)(nil)

func topKFrequent(nums []int, k int) []int {
	freq := make(map[int]int)
	for _, n := range nums {
		freq[n]++
	}

	var h IntCounterHeap
	for num, count := range freq {
		if h.Len() == k && count > h[0].Count {
			heap.Pop(&h)
		}
		if h.Len() < k {
			heap.Push(&h, IntCounter{Val: num, Count: count})
		}

	}

	return h.Vals()
}
