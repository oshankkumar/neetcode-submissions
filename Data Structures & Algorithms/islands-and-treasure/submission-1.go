
type NodePosition [2]int

const Inf = 1<<31 - 1

func islandsAndTreasure(grid [][]int) {
	var queue []NodePosition
	for i, row := range grid {
		for j, cell := range row {
			if cell == 0 {
				queue = append(queue, NodePosition{i, j})
			}
		}
	}

	visited := make(map[NodePosition]struct{})

	var dist int

	dirs := [][2]int{
		{-1, 0}, {1, 0}, {0, -1}, {0, 1},
	}

	for len(queue) > 0 {
		currLen := len(queue)

		for range currLen {
			front := queue[0]
			queue = queue[1:]

			if grid[front[0]][front[1]] == Inf {
				grid[front[0]][front[1]] = dist
			}

			for _, dir := range dirs {
				i, j := front[0]+dir[0], front[1]+dir[1]

				if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) {
					continue
				}

				if grid[i][j] == -1 {
					continue
				}

				nodePos := NodePosition{i, j}

				if _, ok := visited[nodePos]; ok {
					continue
				}

				visited[nodePos] = struct{}{}
				queue = append(queue, nodePos)
			}
		}

		dist++
	}
}
