
func ladderLength(beginWord string, endWord string, wordList []string) int {
	queue := []string{beginWord}
	visited := make(map[string]struct{})
	visited[beginWord] = struct{}{}

	dist := 1
	for len(queue) > 0 {
		currLen := len(queue)

		for range currLen {
			front := queue[0]
			queue = queue[1:]

			if front == endWord {
				return dist
			}

			adjList := adjacentWords(front, wordList)
			for _, word := range adjList {
				if _, ok := visited[word]; ok {
					continue
				}
				visited[word] = struct{}{}
				queue = append(queue, word)
			}

		}

		dist++
	}

	return 0
}

func adjacentWords(start string, wordList []string) []string {
	var result []string
	for _, word := range wordList {
		if wordDistance(start, word) == 1 {
			result = append(result, word)
		}
	}
	return result
}

func wordDistance(w1, w2 string) int {
	var dist int
	for i, b := range []byte(w1) {
		if w2[i] != b {
			dist++
		}
	}
	return dist
}
