
func MaxAreaOfIsland(grid [][]int) int {
	return maxAreaOfIsland(grid)
}

func maxAreaOfIsland(grid [][]int) int {
	visited := make(map[Position]struct{})

	var maxLen int

	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 0 {
				continue
			}
			p := Position{Row: i, Col: j}
			if _, ok := visited[p]; ok {
				continue
			}
			prevLen := len(visited)
			traverseBFS(grid, p, visited)
			maxLen = max(maxLen, len(visited)-prevLen)
		}
	}

	return maxLen
}

type Position struct {
	Row, Col int
}

func (p Position) add(i, j int) Position {
	return Position{Row: p.Row + i, Col: p.Col + j}
}

func (p Position) Left() Position {
	return p.add(0, -1)
}

func (p Position) Right() Position {
	return p.add(0, 1)
}

func (p Position) Up() Position {
	return p.add(-1, 0)
}

func (p Position) Down() Position {
	return p.add(1, 0)
}

func (p Position) AllDir() []Position {
	return []Position{p.Left(), p.Right(), p.Up(), p.Down()}
}

func traverseBFS(grid [][]int, start Position, visted map[Position]struct{}) {
	queue := list.New()
	queue.PushBack(start)

	visted[start] = struct{}{}

	for queue.Len() > 0 {
		p := queue.Remove(queue.Front()).(Position)
		for _, nbr := range p.AllDir() {
			if nbr.Row < 0 || nbr.Row >= len(grid) || nbr.Col < 0 || nbr.Col >= len(grid[0]) {
				continue
			}

			if grid[nbr.Row][nbr.Col] != 1 {
				continue
			}

			if _, ok := visted[nbr]; ok {
				continue
			}

			queue.PushBack(nbr)
			visted[nbr] = struct{}{}
		}
	}
}
