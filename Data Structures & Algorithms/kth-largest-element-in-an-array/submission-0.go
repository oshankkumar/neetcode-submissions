
type MinHeapInt []int

// Len implements [heap.Interface].
func (m MinHeapInt) Len() int {
	return len(m)
}

// Less implements [heap.Interface].
func (m MinHeapInt) Less(i int, j int) bool {
	return m[i] < m[j]
}

// Pop implements [heap.Interface].
func (m *MinHeapInt) Pop() any {
	old, n := *m, len(*m)
	ele := old[n-1]
	*m = old[:n-1]
	return ele
}

// Push implements [heap.Interface].
func (m *MinHeapInt) Push(x any) {
	*m = append(*m, x.(int))
}

// Swap implements [heap.Interface].
func (m MinHeapInt) Swap(i int, j int) {
	m[i], m[j] = m[j], m[i]
}

var _ heap.Interface = (*MinHeapInt)(nil)

func findKthLargest(nums []int, k int) int {
	var minHp MinHeapInt
	for _, n := range nums {
		if minHp.Len() < k {
			heap.Push(&minHp, n)
			continue
		}
		if n > minHp[0] {
			heap.Pop(&minHp)
			heap.Push(&minHp, n)
		}
	}
	return minHp[0]
}