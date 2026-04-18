
type Position struct {
	Row    int
	Column int
}

func (p Position) add(i, j int) Position {
	return Position{
		Row:    p.Row + i,
		Column: p.Column + j,
	}
}

func (p Position) Up() Position {
	return p.add(-1, 0)
}

func (p Position) Down() Position {
	return p.add(1, 0)
}

func (p Position) Left() Position {
	return p.add(0, -1)
}

func (p Position) Right() Position {
	return p.add(0, 1)
}

func (p Position) AllDir() []Position {
	return []Position{p.Left(), p.Right(), p.Up(), p.Down()}
}

func orangesRotting(grid [][]int) int {
	queue := list.New()
	visited := make(map[Position]struct{})

	var totalFruits int
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] > 0 {
				totalFruits++
			}
			if grid[i][j] == 2 {
				p := Position{Row: i, Column: j}
				queue.PushBack(p)
				visited[p] = struct{}{}
			}
		}
	}

	var minDur int
	for len(visited) != totalFruits {
		minDur++
		l := queue.Len()
		if l == 0 {
			return -1
		}

		for range l {
			p := queue.Remove(queue.Front()).(Position)

			for _, nbr := range p.AllDir() {
				if _, ok := visited[nbr]; ok {
					continue
				}

				if nbr.Row < 0 || nbr.Row >= len(grid) || nbr.Column < 0 || nbr.Column >= len(grid[0]) {
					continue
				}
				
				if grid[nbr.Row][nbr.Column] == 1 {
					queue.PushBack(nbr)
					visited[nbr] = struct{}{}
				}
			}
		}
	}

	return minDur
}
