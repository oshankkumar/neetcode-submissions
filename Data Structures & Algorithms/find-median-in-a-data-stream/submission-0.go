
func Constructor() MedianFinder {
    return NewMedianFinder()
}

type IntHeap struct {
	arr      []int
	lessFunc func(arr []int, i int, j int) bool
}

func NewIntHeap(f func(arr []int, i int, j int) bool) *IntHeap {
	return &IntHeap{
		arr:      make([]int, 0),
		lessFunc: f,
	}
}

// Len implements [heap.Interface].
func (h IntHeap) Len() int {
	return len(h.arr)
}

// Less implements [heap.Interface].
func (h IntHeap) Less(i int, j int) bool {
	return h.lessFunc(h.arr, i, j)
}

// Pop implements [heap.Interface].
func (h *IntHeap) Pop() any {
	n := len(h.arr)
	ele := h.arr[n-1]
	h.arr = h.arr[:n-1]
	return ele
}

// Push implements [heap.Interface].
func (h *IntHeap) Push(x any) {
	h.arr = append(h.arr, x.(int))
}

// Swap implements [heap.Interface].
func (h *IntHeap) Swap(i int, j int) {
	h.arr[i], h.arr[j] = h.arr[j], h.arr[i]
}

type MedianFinder struct {
	minHeap *IntHeap
	maxHeap *IntHeap
}

func NewMedianFinder() MedianFinder {
	return MedianFinder{
		minHeap: NewIntHeap(func(arr []int, i, j int) bool {
			return arr[i] < arr[j]
		}),
		maxHeap: NewIntHeap(func(arr []int, i, j int) bool {
			return arr[i] > arr[j]
		}),
	}
}

func (m *MedianFinder) AddNum(num int) {
	heapToPush := m.maxHeap
	if heapToPush.Len() > 0 && num > heapToPush.arr[0] {
		heapToPush = m.minHeap
	}
	heap.Push(heapToPush, num)

	if m.maxHeap.Len() > m.minHeap.Len()+1 {
		num := heap.Pop(m.maxHeap).(int)
		heap.Push(m.minHeap, num)
	}

	if m.minHeap.Len() > m.maxHeap.Len()+1 {
		num := heap.Pop(m.minHeap).(int)
		heap.Push(m.maxHeap, num)
	}
}

func (m *MedianFinder) FindMedian() float64 {
	if m.maxHeap.Len() == m.minHeap.Len() {
		return (float64(m.maxHeap.arr[0]) + float64(m.minHeap.arr[0])) / 2
	}
	heap := m.maxHeap
	if heap.Len() < m.minHeap.Len() {
		heap = m.minHeap
	}
	return float64(heap.arr[0])
}
