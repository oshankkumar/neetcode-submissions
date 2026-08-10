/**
 * Definition of Interval:
 * type Interval struct {
 *    start int
 *    end   int
 * }
 */


type Intervals []Interval

// Len implements [sort.Interface].
func (s Intervals) Len() int {
	return len(s)
}

// Less implements [sort.Interface].
func (s Intervals) Less(i int, j int) bool {
	return s[i].start < s[j].start
}

// Swap implements [sort.Interface].
func (s Intervals) Swap(i int, j int) {
	s[i], s[j] = s[j], s[i]
}

var _ sort.Interface = (Intervals)(nil)

func canAttendMeetings(intervals []Interval) bool {
	if len(intervals) == 0 {
		return true
	}

	sort.Sort(Intervals(intervals))
	prev := intervals[0]
	for _, curr := range intervals[1:] {
		if curr.start < prev.end {
			return false
		}
		prev = curr
	}
	return true
}
