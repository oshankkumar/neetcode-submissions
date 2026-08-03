
type UnionFind struct {
	parent []int
	size   []int
}

func NewUnionFind(n int) *UnionFind {
	uf := &UnionFind{
		parent: make([]int, n),
		size:   make([]int, n),
	}

	for i := range n {
		uf.parent[i] = i
		uf.size[i] = 1
	}

	return uf
}

func (uf *UnionFind) Find(x int) int {
	if uf.parent[x] == x {
		return x
	}

	uf.parent[x] = uf.Find(uf.parent[x])
	return uf.parent[x]
}

func (uf *UnionFind) Union(x, y int) bool {
	px, py := uf.Find(x), uf.Find(y)
	if px == py {
		return false
	}

	if uf.size[px] < uf.size[py] {
		px, py = py, px
	}

	uf.parent[py] = px
	uf.size[px] += uf.size[py]
	return true
}


func findRedundantConnection(edges [][]int) []int {
	uf := NewUnionFind(len(edges) + 1)

	var (
		redundantEdge []int
	)

	for _, edge := range edges {
		if !uf.Union(edge[0], edge[1]) {
			redundantEdge = edge
		}
	}

	return redundantEdge
}