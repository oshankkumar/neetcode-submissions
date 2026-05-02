
type HeapItem struct {
	Row  int
	Col  int
	Diff int
}

type MinHeap []HeapItem

// Len implements [heap.Interface].
func (m MinHeap) Len() int {
	return len(m)
}

// Less implements [heap.Interface].
func (m MinHeap) Less(i int, j int) bool {
	return m[i].Diff < m[j].Diff
}

// Pop implements [heap.Interface].
func (m *MinHeap) Pop() any {
	old := *m
	n := len(old)
	item := old[n-1]
	*m = old[:n-1]
	return item
}

// Push implements [heap.Interface].
func (m *MinHeap) Push(x any) {
	*m = append(*m, x.(HeapItem))
}

// Swap implements [heap.Interface].
func (m MinHeap) Swap(i int, j int) {
	m[i], m[j] = m[j], m[i]
}

var _ heap.Interface = (*MinHeap)(nil)

func minimumEffortPath(heights [][]int) int {
	m, n := len(heights), len(heights[0])
	effort := make([][]int, m)
	for i := 0; i < m; i++ {
		effort[i] = make([]int, n)
		for j := range effort[i] {
			effort[i][j] = 1<<31 - 1
		}
	}

	effort[0][0] = 0

	var minHeap MinHeap
	heap.Push(&minHeap, HeapItem{0, 0, 0})

	dirs := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	for minHeap.Len() > 0 {
		curr := heap.Pop(&minHeap).(HeapItem)
		row, col, diff := curr.Row, curr.Col, curr.Diff

		if effort[row][col] < diff {
			continue
		}

		if row == m-1 && col == n-1 {
			return diff
		}

		for _, dir := range dirs {
			adjR, adjC := row+dir[0], col+dir[1]
			if adjR < 0 || adjR > m-1 || adjC < 0 || adjC > n-1 {
				continue
			}
			newDiff := max(diff, abs(heights[row][col]-heights[adjR][adjC]))
			if newDiff < effort[adjR][adjC] {
				effort[adjR][adjC] = newDiff
				heap.Push(&minHeap, HeapItem{adjR, adjC, newDiff})
			}
		}

	}

	return effort[m-1][n-1]
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
