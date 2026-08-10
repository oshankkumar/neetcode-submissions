/**
 * Definition of Interval:
 * type Interval struct {
 *    start int
 *    end   int
 * }
 */

type Intervals []Interval

// Pop implements [heap.Interface].
func (s *Intervals) Pop() any {
	old := *s
	n := len(old)
	item := old[n-1]
	*s = old[:n-1]
	return item
}

// Push implements [heap.Interface].
func (s *Intervals) Push(x any) {
	*s = append(*s, x.(Interval))
}

// Len implements [sort.Interface].
func (s Intervals) Len() int {
	return len(s)
}

// Less implements [sort.Interface].
func (s Intervals) Less(i int, j int) bool {
	return s[i].end < s[j].end
}

// Swap implements [sort.Interface].
func (s Intervals) Swap(i int, j int) {
	s[i], s[j] = s[j], s[i]
}

var _ heap.Interface = (*Intervals)(nil)

func minMeetingRooms(intervals []Interval) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].start < intervals[j].start
	})

	var h Intervals

	for _, interval := range intervals {
		if h.Len() == 0 {
			heap.Push(&h, interval)
			continue
		}

		top := h[0]

		if top.end <= interval.start {
			heap.Pop(&h)
		}

		heap.Push(&h, interval)
	}

	return h.Len()
}
