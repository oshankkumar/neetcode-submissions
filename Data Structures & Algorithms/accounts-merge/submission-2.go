type UnionFind struct {
	parent []int
	rank   []int
}

func NewUnionFind(n int) *UnionFind {
	u := &UnionFind{
		parent: make([]int, n),
		rank:   make([]int, n),
	}
	for i := range n {
		u.parent[i] = i
		u.rank[i] = 1
	}
	return u
}

func (u *UnionFind) Find(x int) int {
	if u.parent[x] != x {
		u.parent[x] = u.Find(u.parent[x])
	}
	return u.parent[x]
}

func (u *UnionFind) Union(x, y int) bool {
	pX, pY := u.Find(x), u.Find(y)
	if pX == pY {
		return false
	}

	rX, rY := u.rank[x], u.rank[y]

	if rX == rY {
		u.parent[pY] = pX
		u.rank[pX]++
	} else if rX > rY {
		u.parent[pY] = pX
	} else {
		u.parent[pX] = pY
	}

	return true
}

func accountsMerge(accounts [][]string) [][]string {
	emailsToAccID := make(map[string]int)

	uf := NewUnionFind(len(accounts))

	for accID, acc := range accounts {
		for _, email := range acc[1:] {
			idx, ok := emailsToAccID[email]
			if !ok {
				emailsToAccID[email] = accID
				continue
			}
			uf.Union(accID, idx)
		}
	}

	emailGroups := make(map[int][]string)
	for email, accID := range emailsToAccID {
		root := uf.Find(accID)
		emailGroups[root] = append(emailGroups[root], email)
	}

	var results [][]string
	for accID, emails := range emailGroups {
		sort.Strings(emails)
		result := append([]string{accounts[accID][0]}, emails...)
		results = append(results, result)
	}
	return results
}
