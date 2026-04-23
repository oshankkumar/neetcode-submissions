
type CacheItem struct {
	Key        int
	Val        int
	index      int
	freq       int
	accessedAt int
}

var _ heap.Interface = (*MinHeap)(nil)

type MinHeap []*CacheItem

// Len implements [heap.Interface].
func (m *MinHeap) Len() int {
	return len(*m)
}

// Less implements [heap.Interface].
func (m *MinHeap) Less(i int, j int) bool {
	minHeap := *m
	if minHeap[i].freq == minHeap[j].freq {
		return minHeap[i].accessedAt < minHeap[j].accessedAt
	}
	return minHeap[i].freq < minHeap[j].freq
}

// Pop implements [heap.Interface].
func (m *MinHeap) Pop() any {
	old, n := *m, len(*m)
	item := old[n-1]
	item.index = -1
	old[n-1] = nil
	*m = old[:n-1]
	return item
}

// Push implements [heap.Interface].
func (m *MinHeap) Push(x any) {
	c := x.(*CacheItem)
	c.index = m.Len()
	*m = append(*m, c)
}

// Swap implements [heap.Interface].
func (m *MinHeap) Swap(i int, j int) {
	heap := *m
	heap[i], heap[j] = heap[j], heap[i]
	heap[i].index = i
	heap[j].index = j
}

type LFUCache struct {
	lookupTable map[int]*CacheItem
	itemHeap    *MinHeap
	capacity    int
	currTs      int
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		lookupTable: make(map[int]*CacheItem),
		itemHeap:    new(MinHeap),
		capacity:    capacity,
	}
}

func (l *LFUCache) Get(key int) int {
	l.currTs++

	item, ok := l.lookupTable[key]
	if !ok {
		return -1
	}

	item.freq++
	item.accessedAt = l.currTs
	heap.Fix(l.itemHeap, item.index)
	return item.Val
}

func (l *LFUCache) Put(key int, value int) {
	l.currTs++

	if item, ok := l.lookupTable[key]; ok {
		item.Val = value
		item.freq++
		item.accessedAt = l.currTs
		heap.Fix(l.itemHeap, item.index)
		return
	}

	if l.itemHeap.Len() == l.capacity {
		item := heap.Pop(l.itemHeap).(*CacheItem)
		delete(l.lookupTable, item.Key)
	}

	item := &CacheItem{Key: key, Val: value, accessedAt: l.currTs, freq: 1}
	heap.Push(l.itemHeap, item)
	l.lookupTable[key] = item
}
