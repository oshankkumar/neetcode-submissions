func Constructor(capacity int) LRUCache {
    return NewLRUCache(capacity)
}

type cacheItem struct {
	val int
	key int
}

type LRUCache struct {
	cap   int
	list  *list.List
	store map[int]*list.Element
}

func NewLRUCache(capacity int) LRUCache {
	return LRUCache{
		cap:   capacity,
		list:  list.New(),
		store: make(map[int]*list.Element),
	}
}

func (l *LRUCache) Get(key int) int {
	node, ok := l.store[key]
	if !ok {
		return -1
	}

	l.list.MoveToFront(node)
	return node.Value.(cacheItem).val
}

func (l *LRUCache) Put(key int, value int) {
	if node, ok := l.store[key]; ok {
		l.list.MoveToFront(node)
		node.Value = cacheItem{key: key, val: value}
		return
	}

	if len(l.store) < l.cap {
		l.store[key] = l.list.PushFront(cacheItem{key: key, val: value})
		return
	}

	delItem := l.list.Remove(l.list.Back()).(cacheItem)
	delete(l.store, delItem.key)
	l.store[key] = l.list.PushFront(cacheItem{key: key, val: value})
}
