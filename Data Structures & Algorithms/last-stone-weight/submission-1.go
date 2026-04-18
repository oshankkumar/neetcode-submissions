
var _ heap.Interface = (*MaxHeapStone)(nil)

type MaxHeapStone []int

// Len implements heap.Interface.
func (m *MaxHeapStone) Len() int {
	return len(*m)
}

// Less implements heap.Interface.
func (m *MaxHeapStone) Less(i int, j int) bool {
	return (*m)[i] > (*m)[j]
}

// Pop implements heap.Interface.
func (m *MaxHeapStone) Pop() any {
	old, n := *m, len(*m)
	e := old[n-1]
	*m = old[:n-1]
	return e
}

// Push implements heap.Interface.
func (m *MaxHeapStone) Push(x any) {
	*m = append(*m, x.(int))
}

// Swap implements heap.Interface.
func (m *MaxHeapStone) Swap(i int, j int) {
	(*m)[i], (*m)[j] = (*m)[j], (*m)[i]
}

func lastStoneWeight(stones []int) int {
	h := MaxHeapStone(stones)
	heap.Init(&h)
	for h.Len() > 1 {
		l1, l2 := heap.Pop(&h).(int), heap.Pop(&h).(int)
		heap.Push(&h, abs(l1-l2))
	}
	return heap.Pop(&h).(int)
}

func abs(n int) int {
	if n >= 0 {
		return n
	}
	return -n
}
