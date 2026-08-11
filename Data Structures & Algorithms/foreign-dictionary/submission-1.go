func alienOrder(words []string) string {
	// Every character must exist in the graph,
	// even if it has no edges.
	graph := make(map[byte][]byte)
	indegree := make(map[byte]int)

	for _, word := range words {
		for i := 0; i < len(word); i++ {
			c := word[i]
			if _, exists := indegree[c]; !exists {
				indegree[c] = 0
			}
		}
	}

	// Build ordering constraints.
	for i := 0; i < len(words)-1; i++ {
		w1 := words[i]
		w2 := words[i+1]

		minLen := min(len(w1), len(w2))

		foundDifference := false

		for j := 0; j < minLen; j++ {
			if w1[j] != w2[j] {
				from := w1[j]
				to := w2[j]

				// Avoid duplicate edges.
				if !hasEdge(graph[from], to) {
					graph[from] = append(graph[from], to)
					indegree[to]++
				}

				foundDifference = true
				break
			}
		}

		// Invalid:
		// ["abc", "ab"]
		if !foundDifference && len(w1) > len(w2) {
			return ""
		}
	}

	// Kahn's algorithm.
	queue := make([]byte, 0)

	for c, degree := range indegree {
		if degree == 0 {
			queue = append(queue, c)
		}
	}

	result := make([]byte, 0, len(indegree))

	for len(queue) > 0 {
		c := queue[0]
		queue = queue[1:]

		result = append(result, c)

		for _, next := range graph[c] {
			indegree[next]--

			if indegree[next] == 0 {
				queue = append(queue, next)
			}
		}
	}

	// Not all characters were processed => cycle.
	if len(result) != len(indegree) {
		return ""
	}

	return string(result)
}

func hasEdge(edges []byte, target byte) bool {
	for _, c := range edges {
		if c == target {
			return true
		}
	}
	return false
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func foreignDictionary(words []string) string {
	return alienOrder(words)
}
