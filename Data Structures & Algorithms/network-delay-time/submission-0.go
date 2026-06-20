
type NetworkNodeID int

type NetworkEdge struct {
	To     NetworkNodeID
	Weight int
}

type pqItem struct {
	node     NetworkNodeID
	distance int
}

var _ heap.Interface = (*priorityQueue)(nil)

type priorityQueue []pqItem

// Len implements [heap.Interface].
func (p *priorityQueue) Len() int {
	return len(*p)
}

// Less implements [heap.Interface].
func (p priorityQueue) Less(i int, j int) bool {
	return p[i].distance < p[j].distance
}

// Pop implements [heap.Interface].
func (p *priorityQueue) Pop() any {
	old := *p
	n := len(old)
	item := old[n-1]
	*p = old[:n-1]
	return item
}

// Push implements [heap.Interface].
func (p *priorityQueue) Push(x any) {
	*p = append(*p, x.(pqItem))
}

// Swap implements [heap.Interface].
func (p priorityQueue) Swap(i int, j int) {
	p[i], p[j] = p[j], p[i]
}

func networkDelayTime(times [][]int, n int, k int) int {
	graph := make(map[NetworkNodeID][]NetworkEdge)
	src := NetworkNodeID(k)

	for _, time := range times {
		u, v, w := NetworkNodeID(time[0]), NetworkNodeID(time[1]), time[2]
		graph[u] = append(graph[u], NetworkEdge{
			To:     v,
			Weight: w,
		})
	}

	dist := make(map[NetworkNodeID]int)
	for i := range n {
		dist[NetworkNodeID(i+1)] = math.MaxInt
	}
	dist[src] = 0

	var pq priorityQueue
	heap.Push(&pq, pqItem{
		node:     src,
		distance: dist[src],
	})

	for pq.Len() > 0 {
		currNode := heap.Pop(&pq).(pqItem)
		u := currNode.node
		if dist[u] < currNode.distance {
			continue
		}

		for _, edge := range graph[u] {
			v, w := edge.To, edge.Weight
			newDist := dist[u] + w
			if dist[v] > newDist {
				dist[v] = newDist
				heap.Push(&pq, pqItem{
					node:     v,
					distance: dist[v],
				})
			}
		}
	}

	var minDist int
	for _, val := range dist {
		if val == math.MaxInt {
			return -1
		}
		minDist = max(minDist, val)
	}

	return minDist
}
