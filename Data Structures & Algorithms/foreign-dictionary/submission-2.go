
type LetterGraph struct {
	adj      map[byte][]byte
	inDegree map[byte]int
}

func NewLetterGraph() *LetterGraph {
	return &LetterGraph{
		adj:      make(map[byte][]byte),
		inDegree: make(map[byte]int),
	}
}

func (l *LetterGraph) AddEdge(u, v byte) {
	l.adj[u] = append(l.adj[u], v)
	l.inDegree[v]++
}

func ForeignDictionary(words []string) string {
	return foreignDictionary(words)
}

func foreignDictionary(words []string) string {
	graph := NewLetterGraph()

	for _, word := range words {
		for i := 0; i < len(word); i++ {
			graph.inDegree[word[i]] = 0
		}
	}

	for i := 1; i < len(words); i++ {
		prev, curr := words[i-1], words[i]
		if curr == prev {
			continue
		}

		var j int
		for j < min(len(prev), len(curr)) && prev[j] == curr[j] {
			j++
		}

		if j == len(curr) {
			return ""
		}

		if j == len(prev) {
			continue
		}

		graph.AddEdge(prev[j], curr[j])
	}

	var queue []byte

	for l, degree := range graph.inDegree {
		if degree == 0 {
			queue = append(queue, l)
		}
	}

	var result []byte
	for len(queue) > 0 {
		frontLetter := queue[0]
		queue = queue[1:]

		result = append(result, frontLetter)

		for _, adjLetter := range graph.adj[frontLetter] {
			graph.inDegree[adjLetter]--
			if graph.inDegree[adjLetter] == 0 {
				queue = append(queue, adjLetter)
			}
		}
	}

	if len(result) != len(graph.inDegree) {
		return ""
	}

	return string(result)
}
