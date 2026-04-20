import (
	"iter"
	"maps"
)

type NodeID int

func NewCourseGraph() *CourseGraph {
	return &CourseGraph{
		adj:      make(map[NodeID]map[NodeID]struct{}),
		inDegree: make(map[NodeID]int),
	}
}

type CourseGraph struct {
	adj      map[NodeID]map[NodeID]struct{}
	inDegree map[NodeID]int
}

func (g *CourseGraph) AddEdge(from, to NodeID) {
	if g.adj[from] == nil {
		g.adj[from] = make(map[NodeID]struct{})
	}

	if g.adj[to] == nil {
		g.adj[to] = make(map[NodeID]struct{})
	}

	g.adj[from][to] = struct{}{}
	g.inDegree[to]++
}

func (g *CourseGraph) Neighbours(nodeID NodeID) iter.Seq[NodeID] {
	return maps.Keys(g.adj[nodeID])
}

func (g *CourseGraph) InDegree() map[NodeID]int {
	return g.inDegree
}

func checkIfPrerequisite(numCourses int, prerequisites [][]int, queries [][]int) []bool {
	g := NewCourseGraph()

	for _, p := range prerequisites {
		g.AddEdge(NodeID(p[0]), NodeID(p[1]))
	}

	var result []bool

	for _, q := range queries {
		var stack []NodeID
		visited := make(map[NodeID]struct{})

		stack = append(stack, NodeID(q[0]))
		var found bool

		for len(stack) > 0 {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			if top == NodeID(q[1]) {
				found = true
				break
			}

			for n := range g.Neighbours(top) {
				if _, ok := visited[n]; ok {
					continue
				}
				visited[n] = struct{}{}
				stack = append(stack, n)
			}
		}

		result = append(result, found)

	}

	return result
}
