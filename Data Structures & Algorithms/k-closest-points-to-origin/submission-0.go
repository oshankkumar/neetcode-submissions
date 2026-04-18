
type Point [2]int

var _ heap.Interface = (*MaxPointHeap)(nil)

type MaxPointHeap []Point

// Len implements heap.Interface.
func (m MaxPointHeap) Len() int {
	return len(m)
}

// Less implements heap.Interface.
func (m MaxPointHeap) Less(i int, j int) bool {
	return DistFromOrigin(m[i]) > DistFromOrigin(m[j])
}

// Pop implements heap.Interface.
func (m *MaxPointHeap) Pop() any {
	old, n := *m, len(*m)
	e := old[n-1]
	*m = old[:n-1]
	return e
}

// Push implements heap.Interface.
func (m *MaxPointHeap) Push(x any) {
	*m = append(*m, x.(Point))
}

// Swap implements heap.Interface.
func (m MaxPointHeap) Swap(i int, j int) {
	m[i], m[j] = m[j], m[i]
}

func kClosest(points [][]int, k int) [][]int {
	var h MaxPointHeap

	for _, p := range points {
		point := Point(p)
		if h.Len() < k {
			heap.Push(&h, point)
			continue
		}
		if DistFromOrigin(point) < DistFromOrigin(h[0]) {
			heap.Pop(&h)
			heap.Push(&h, point)
		}
	}

	var result [][]int
	for _, p := range h {
		result = append(result, p[:])
	}
	return result
}

func DistFromOrigin(p Point) float64 {
	return math.Sqrt(float64(p[0]*p[0] + p[1]*p[1]))
}
