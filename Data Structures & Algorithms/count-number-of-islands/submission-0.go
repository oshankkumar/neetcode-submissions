
func NewLandGrid(grid [][]byte) *LandGrid {
	return &LandGrid{grid: grid}
}

type LandGrid struct {
	grid [][]byte
}

type Coordinate struct {
	X, Y int
}

func (l *LandGrid) DisjointCount() int {
	visited := make(map[Coordinate]struct{})
	var count int
	for i := range l.grid {
		for j := range l.grid[i] {
			if l.grid[i][j] == '0' {
				continue
			}
			c := Coordinate{i, j}
			if _, ok := visited[c]; ok {
				continue
			}
			l.traverse(c, visited)
			count++
		}

	}
	return count
}

func (l *LandGrid) traverse(start Coordinate, visted map[Coordinate]struct{}) {
	queue := list.New()
	queue.PushBack(start)

	validPoint := func(c Coordinate) bool {
		return c.X >= 0 && c.X < len(l.grid) && c.Y >= 0 && c.Y < len(l.grid[0])
	}

	for queue.Len() > 0 {
		fc := queue.Remove(queue.Front()).(Coordinate)

		for _, c := range []Coordinate{
			{fc.X - 1, fc.Y},
			{fc.X + 1, fc.Y},
			{fc.X, fc.Y - 1},
			{fc.X, fc.Y + 1},
		} {
			if !validPoint(c) {
				continue
			}

			if l.grid[c.X][c.Y] == '0' {
				continue
			}
			
			if _, ok := visted[c]; ok {
				continue
			}
			queue.PushBack(c)
			visted[c] = struct{}{}
		}
	}
}

func numIslands(grid [][]byte) int {
	return NewLandGrid(grid).DisjointCount()
}
