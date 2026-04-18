func Constructor(capacity int) LRUCache {
    return NewLRUCache(capacity)
}

type cacheItem struct {
	val int
	key *list.Element
}

type LRUCache struct {
	cap   int
	list  *list.List
	store map[int]*cacheItem
}

func NewLRUCache(capacity int) LRUCache {
	return LRUCache{
		cap:   capacity,
		list:  list.New(),
		store: make(map[int]*cacheItem),
	}
}

func (l *LRUCache) Get(key int) int {
	item, ok := l.store[key]
	if !ok {
		return -1
	}
	l.list.MoveToFront(item.key)
	return item.val
}

func (l *LRUCache) Put(key int, value int) {
	if item, ok := l.store[key]; ok {
		item.val = value
		l.list.MoveToFront(item.key)
		return
	}

	if len(l.store) < l.cap {
		l.store[key] = &cacheItem{val: value, key: l.list.PushFront(key)}
		return
	}

	delKey := l.list.Remove(l.list.Back()).(int)
	delete(l.store, delKey)

	l.store[key] = &cacheItem{val: value, key: l.list.PushFront(key)}
}
