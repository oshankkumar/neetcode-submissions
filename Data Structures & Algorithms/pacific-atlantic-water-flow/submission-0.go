
type IslandPosition struct {
	Row    int
	Column int
}

func (p IslandPosition) AllDirs(m, n int) []IslandPosition {
	var all []IslandPosition
	dirs := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for _, d := range dirs {
		i, j := p.Row+d[0], p.Column+d[1]
		if i >= 0 && i < m && j >= 0 && j < n {
			all = append(all, IslandPosition{Row: i, Column: j})
		}
	}
	return all
}

func pacificAtlantic(heights [][]int) [][]int {
	var i, j int
	visitedPacific := make(map[IslandPosition]struct{})

	i = 0
	for j = 0; j < len(heights[i]); j++ {
		p := IslandPosition{Row: i, Column: j}
		if _, ok := visitedPacific[p]; ok {
			continue
		}
		visitedPacific[p] = struct{}{}
		traverseIslandsDFS(p, heights, visitedPacific)
	}

	j = 0
	for i = 0; i < len(heights); i++ {
		p := IslandPosition{Row: i, Column: j}
		if _, ok := visitedPacific[p]; ok {
			continue
		}
		visitedPacific[p] = struct{}{}
		traverseIslandsDFS(p, heights, visitedPacific)
	}

	visitedAtlantic := make(map[IslandPosition]struct{})

	i = len(heights) - 1
	for j = 0; j < len(heights[i]); j++ {
		p := IslandPosition{Row: i, Column: j}
		if _, ok := visitedAtlantic[p]; ok {
			continue
		}
		visitedAtlantic[p] = struct{}{}
		traverseIslandsDFS(p, heights, visitedAtlantic)
	}

	j = len(heights[0]) - 1
	for i = 0; i < len(heights); i++ {
		p := IslandPosition{Row: i, Column: j}
		if _, ok := visitedAtlantic[p]; ok {
			continue
		}
		visitedAtlantic[p] = struct{}{}
		traverseIslandsDFS(p, heights, visitedAtlantic)
	}

	visitedBoth := make([][]int, 0)
	for p := range visitedAtlantic {
		if _, ok := visitedPacific[p]; ok {
			visitedBoth = append(visitedBoth, []int{p.Row, p.Column})
		}
	}
	return visitedBoth
}

func traverseIslandsDFS(start IslandPosition, heights [][]int, visited map[IslandPosition]struct{}) {
	for _, n := range start.AllDirs(len(heights), len(heights[0])) {
		if _, ok := visited[n]; ok {
			continue
		}
		if heights[n.Row][n.Column] < heights[start.Row][start.Column] {
			continue
		}
		visited[n] = struct{}{}
		traverseIslandsDFS(n, heights, visited)
	}
}
