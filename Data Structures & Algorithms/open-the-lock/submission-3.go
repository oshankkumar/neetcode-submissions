
type lockNode struct {
	key  string
	move int
}

func openLock(deadends []string, target string) int {
	deps := make(map[string]struct{})
	for _, d := range deadends {
		deps[d] = struct{}{}
	}

	q := make([]*lockNode, 0)
	head := 0

	if _, ok := deps["0000"]; ok {
		return -1
	}

	q = append(q, &lockNode{key: "0000", move: 0})
	visited := make(map[string]struct{})
	visited["0000"] = struct{}{}

	for len(q) > head {
		node := q[head]
		head++

		if node.key == target {
			return node.move
		}

		for _, child := range allPossibleMoves(node.key) {
			if _, ok := visited[child]; ok {
				continue
			}
			if _, ok := deps[child]; ok {
				continue
			}
			q = append(q, &lockNode{key: child, move: node.move + 1})
			visited[child] = struct{}{}
		}
	}

	return -1
}

func allPossibleMoves(target string) []string {
	result := make([]string, 0)
	for i, r := range target {
		num1, num2 := []rune(target), []rune(target)
		digit := r - '0'
		next := (digit + 1) % 10
		prev := (digit - 1 + 10) % 10
		num1[i] = '0' + rune(next)
		num2[i] = '0' + rune(prev)
		_ = r
		result = append(result, string(num1), string(num2))
	}
	return result
}
