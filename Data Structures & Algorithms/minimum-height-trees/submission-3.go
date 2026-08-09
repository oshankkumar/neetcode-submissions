
type MHTGraph struct {
	AdjLists [][]int
	Degrees  []int
}

func (g *MHTGraph) AddEdge(u, v int) {
	g.AdjLists[u] = append(g.AdjLists[u], v)
	g.AdjLists[v] = append(g.AdjLists[v], u)
	g.Degrees[u]++
	g.Degrees[v]++
}

func NewMHTGraph(n int, edges [][]int) *MHTGraph {
	g := &MHTGraph{
		AdjLists: make([][]int, n),
		Degrees:  make([]int, n),
	}

	for _, edge := range edges {
		g.AddEdge(edge[0], edge[1])
	}

	return g
}

func findMinHeightTrees(n int, edges [][]int) []int {
	g := NewMHTGraph(n, edges)

	var queue []int
	for i, degree := range g.Degrees {
		if degree <= 1 {
			queue = append(queue, i)
		}
	}

	for len(queue) > 0 {
		if n <= 2 {
			return queue
		}

		currLen := len(queue)

		for range currLen {
			node := queue[0]
			queue = queue[1:]
			n--

			for _, adj := range g.AdjLists[node] {
				g.Degrees[adj]--
				if g.Degrees[adj] == 1 {
					queue = append(queue, adj)
				}
			}

		}
	}

	return nil
}
