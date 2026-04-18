func Constructor() TimeMap {
	return NewTimeMap()
}

type MapItem struct {
	Value string
	Ts    int
}

type MapItems []MapItem

type TimeMap struct {
	store map[string]*MapItems
}

func NewTimeMap() TimeMap {
	return TimeMap{store: make(map[string]*MapItems)}
}

func (t *TimeMap) Set(key string, value string, timestamp int) {
	items, ok := t.store[key]
	if !ok {
		items = &MapItems{}
		t.store[key] = items
	}

	i := sort.Search(len(*items), func(i int) bool {
		return (*items)[i].Ts <= timestamp
	})

	if i == len(*items) {
		*items = append(*items, MapItem{Value: value, Ts: timestamp})
		return
	}

	left := (*items)[0:i]
	right := (*items)[i:]

	arr := make([]MapItem, len(*items)+1)

	copy(arr[0:len(left)], left)
	arr[len(left)] = MapItem{Value: value, Ts: timestamp}
	copy(arr[len(left)+1:], right)

	*items = arr
}

func (t *TimeMap) Get(key string, timestamp int) string {
	items, ok := t.store[key]
	if !ok {
		return ""
	}

	i := sort.Search(len(*items), func(i int) bool {
		return (*items)[i].Ts <= timestamp
	})
	if i < len(*items) {
		return (*items)[i].Value
	}
	return ""
}
