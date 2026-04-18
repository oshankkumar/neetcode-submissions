
func FindOrder(numCourses int, prerequisites [][]int) []int {
	return findOrder(numCourses, prerequisites)
}

type Vertex struct {
	Key     int
	AdjList []*Vertex
	degree  int
}

func NewGraph(keyCount int) *Graph {
	g := &Graph{Nodes: make(map[int]*Vertex)}
	for key := range keyCount {
		g.Nodes[key] = &Vertex{Key: key}
	}
	return g
}

type Graph struct {
	Nodes map[int]*Vertex
}

func (g *Graph) AddEdge(key1, key2 int) {
	g.Nodes[key1].AdjList = append(g.Nodes[key1].AdjList, g.Nodes[key2])
	g.Nodes[key2].degree++
}

func findOrder(numCourses int, prerequisites [][]int) []int {
	g := NewGraph(numCourses)

	for _, p := range prerequisites {
		g.AddEdge(p[1], p[0])
	}

	var q []*Vertex

	for _, v := range g.Nodes {
		if v.degree == 0 {
			q = append(q, v)
		}
	}

	var result []int

	for len(q) > 0 {
		front := q[0]
		q = q[1:]

		result = append(result, front.Key)

		for _, v := range front.AdjList {
			v.degree--
			if v.degree == 0 {
				q = append(q, v)
			}
		}
	}

	if len(result) == numCourses {
		return result
	}
	return nil
}
