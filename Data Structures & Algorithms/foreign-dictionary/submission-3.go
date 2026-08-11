
type LetterGraph struct {
	Adj      map[byte][]byte
	InDegree map[byte]int
}

func NewLetterGraph() *LetterGraph {
	return &LetterGraph{
		Adj:      make(map[byte][]byte),
		InDegree: make(map[byte]int),
	}
}

func (l *LetterGraph) AddNode(u byte) {
	if _, ok := l.Adj[u]; ok {
		return
	}
	l.Adj[u] = make([]byte, 0)
	l.InDegree[u] = 0
}

func (l *LetterGraph) AddEdge(u, v byte) {
	l.Adj[u] = append(l.Adj[u], v)
	l.InDegree[v]++
}

func ForeignDictionary(words []string) string {
	return foreignDictionary(words)
}

func foreignDictionary(words []string) string {
	graph := NewLetterGraph()

	for _, word := range words {
		for i := 0; i < len(word); i++ {
			graph.AddNode(word[i])
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

	for l, degree := range graph.InDegree {
		if degree == 0 {
			queue = append(queue, l)
		}
	}

	var result []byte
	for len(queue) > 0 {
		frontLetter := queue[0]
		queue = queue[1:]

		result = append(result, frontLetter)

		for _, adjLetter := range graph.Adj[frontLetter] {
			graph.InDegree[adjLetter]--
			if graph.InDegree[adjLetter] == 0 {
				queue = append(queue, adjLetter)
			}
		}
	}

	if len(result) != len(graph.InDegree) {
		return ""
	}

	return string(result)
}
