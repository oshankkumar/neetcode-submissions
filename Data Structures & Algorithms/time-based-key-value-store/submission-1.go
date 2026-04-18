func Constructor() TimeMap {
   return NewTimeMap()
}


type MapItem struct {
	Value string
	Ts    int
}

type TimeMap struct {
	store map[string][]MapItem
}

func NewTimeMap() TimeMap {
	return TimeMap{store: make(map[string][]MapItem)}
}

func (t *TimeMap) Set(key string, value string, timestamp int) {
	t.store[key] = append(t.store[key], MapItem{Value: value, Ts: timestamp})
}

func (t *TimeMap) Get(key string, timestamp int) string {
	items, ok := t.store[key]
	if !ok {
		return ""
	}

	i := sort.Search(len(items), func(i int) bool {
		return items[i].Ts > timestamp
	})

	if i == 0 {
		return ""
	}

	return items[i-1].Value
}
