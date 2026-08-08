
func solve(board [][]byte) {
	visited := make(map[[2]int]struct{})

	idxs := []int{0, len(board) - 1}
	for _, i := range idxs {
		for j := range len(board[0]) {
			node := [2]int{i, j}
			_, ok := visited[node]
			if board[i][j] == 'O' && !ok {
				dfsBoard(board, node, visited)
			}
		}
	}

	idxs = []int{0, len(board[0]) - 1}
	for _, j := range idxs {
		for i := range len(board) {
			node := [2]int{i, j}
			_, ok := visited[node]
			if board[i][j] == 'O' && !ok {
				dfsBoard(board, node, visited)
			}
		}
	}

	for i := range board {
		for j := range board[i] {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
		}
	}

	for i := range board {
		for j := range board[i] {
			if board[i][j] == '#' {
				board[i][j] = 'O'
			}
		}
	}
}

func dfsBoard(board [][]byte, src [2]int, visited map[[2]int]struct{}) {
	dirs := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	board[src[0]][src[1]] = '#'
	visited[src] = struct{}{}

	for _, dir := range dirs {
		i, j := src[0]+dir[0], src[1]+dir[1]
		if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) {
			continue
		}

		if board[i][j] != 'O' {
			continue
		}

		node := [2]int{i, j}

		if _, ok := visited[node]; ok {
			continue
		}

		dfsBoard(board, node, visited)
	}
}
