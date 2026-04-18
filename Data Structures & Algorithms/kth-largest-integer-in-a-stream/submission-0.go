
type MinHeap []int

// Len implements heap.Interface.
func (m *MinHeap) Len() int {
	return len(*m)
}

// Less implements heap.Interface.
func (m *MinHeap) Less(i int, j int) bool {
	return (*m)[i] < (*m)[j]
}

// Pop implements heap.Interface.
func (m *MinHeap) Pop() any {
	old := *m
	e := old[len(old)-1]
	old = old[:len(old)-1]
	*m = old
	return e
}

// Push implements heap.Interface.
func (m *MinHeap) Push(x any) {
	*m = append(*m, x.(int))
}

// Swap implements heap.Interface.
func (m *MinHeap) Swap(i int, j int) {
	(*m)[i], (*m)[j] = (*m)[j], (*m)[i]
}

type KthLargest struct {
	k    int
	nums *MinHeap
}

func Constructor(k int, nums []int) KthLargest {
	var m MinHeap
	kl := KthLargest{k: k, nums: &m}
	for _, num := range nums {
		kl.Add(num)
	}
	return kl
}


func (k *KthLargest) Add(val int) int {
	if k.nums.Len() < k.k {
		heap.Push(k.nums, val)
		return (*k.nums)[0]
	}

	if val > (*k.nums)[0] {
		(*k.nums)[0] = val
		heap.Fix(k.nums, 0)
	}

	return (*k.nums)[0]
}
