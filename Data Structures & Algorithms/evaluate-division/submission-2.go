type Variable string

type Edge struct {
	To    Variable
	Value float64
}

type Graph map[Variable][]Edge

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	graph := make(Graph)

	for i, eq := range equations {
		u := Variable(eq[0])
		v := Variable(eq[1])

		graph[u] = append(graph[u], Edge{
			To:    v,
			Value: values[i],
		})

		graph[v] = append(graph[v], Edge{
			To:    u,
			Value: 1.0 / values[i],
		})
	}

	ans := make([]float64, 0, len(queries))

	for _, q := range queries {
		src := Variable(q[0])
		dst := Variable(q[1])

		// Variable doesn't exist.
		if _, ok := graph[src]; !ok {
			ans = append(ans, -1.0)
			continue
		}

		if _, ok := graph[dst]; !ok {
			ans = append(ans, -1.0)
			continue
		}

		// Same variable.
		if src == dst {
			ans = append(ans, 1.0)
			continue
		}

		visited := make(map[Variable]bool)

		if value, ok := dfs(graph, src, dst, 1.0, visited); ok {
			ans = append(ans, value)
		} else {
			ans = append(ans, -1.0)
		}
	}

	return ans
}

func dfs(
	graph Graph,
	curr Variable,
	target Variable,
	product float64,
	visited map[Variable]bool,
) (float64, bool) {

	if curr == target {
		return product, true
	}

	visited[curr] = true

	for _, edge := range graph[curr] {
		if visited[edge.To] {
			continue
		}

		if value, ok := dfs(
			graph,
			edge.To,
			target,
			product*edge.Value,
			visited,
		); ok {
			return value, true
		}
	}

	return 0, false
}