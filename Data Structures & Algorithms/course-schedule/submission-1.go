
func canFinish(numCourses int, prerequisites [][]int) bool {
	graph := make([][]int, numCourses)
	inDegree := make([]int, numCourses)

	for _, p := range prerequisites {
		graph[p[1]] = append(graph[p[1]], p[0])
		inDegree[p[0]]++
	}

	var q []int

	for i, n := range inDegree {
		if n == 0 {
			q = append(q, i)
		}
	}

	var results []int
	for len(q) > 0 {
		front := q[0]
		q = q[1:]

		results = append(results, front)

		for _, neigh := range graph[front] {
			inDegree[neigh]--
			if inDegree[neigh] == 0 {
				q = append(q, neigh)
			}
		}

	}

	return len(results) == numCourses
}
