
type MeetingRoomLease struct {
	RoomNumber int
	End        int
}

type MeetingRoomLeasePQ []*MeetingRoomLease

var _ heap.Interface = (*MeetingRoomLeasePQ)(nil)

// Len implements [heap.Interface].
func (m MeetingRoomLeasePQ) Len() int {
	return len(m)
}

// Less implements [heap.Interface].
func (m MeetingRoomLeasePQ) Less(i int, j int) bool {
	if m[i].End == m[j].End {
		return m[i].RoomNumber < m[j].RoomNumber
	}
	return m[i].End < m[j].End
}

// Pop implements [heap.Interface].
func (m *MeetingRoomLeasePQ) Pop() any {
	old := *m
	n := len(old)
	item := old[n-1]
	*m = old[:n-1]
	return item
}

// Push implements [heap.Interface].
func (m *MeetingRoomLeasePQ) Push(x any) {
	*m = append(*m, x.(*MeetingRoomLease))
}

// Swap implements [heap.Interface].
func (m MeetingRoomLeasePQ) Swap(i int, j int) {
	m[i], m[j] = m[j], m[i]
}

type IntMinHeap []int

// Len implements [heap.Interface].
func (h IntMinHeap) Len() int {
	return len(h)
}

// Less implements [heap.Interface].
func (h IntMinHeap) Less(i int, j int) bool {
	return h[i] < h[j]
}

// Pop implements [heap.Interface].
func (h *IntMinHeap) Pop() any {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[:n-1]
	return item
}

// Push implements [heap.Interface].
func (h *IntMinHeap) Push(x any) {
	*h = append(*h, x.(int))
}

// Swap implements [heap.Interface].
func (h IntMinHeap) Swap(i int, j int) {
	h[i], h[j] = h[j], h[i]
}

var _ heap.Interface = (*IntMinHeap)(nil)

func mostBooked(n int, meetings [][]int) int {
	sort.Slice(meetings, func(i, j int) bool {
		return meetings[i][0] < meetings[j][0]
	})

	var (
		meetingRoomLeases     MeetingRoomLeasePQ
		availableMeetingRooms IntMinHeap
		bookingCount          = make(map[int]int)
	)

	for i := range n {
		heap.Push(&availableMeetingRooms, i)
	}

	for _, meeting := range meetings {
		for meetingRoomLeases.Len() > 0 && meetingRoomLeases[0].End <= meeting[0] {
			room := heap.Pop(&meetingRoomLeases).(*MeetingRoomLease)
			heap.Push(&availableMeetingRooms, room.RoomNumber)
		}

		if availableMeetingRooms.Len() == 0 {
			earliestFreeLease := heap.Pop(&meetingRoomLeases).(*MeetingRoomLease)
			meetingDur := meeting[1] - meeting[0]

			heap.Push(&meetingRoomLeases, &MeetingRoomLease{
				RoomNumber: earliestFreeLease.RoomNumber,
				End:        earliestFreeLease.End + meetingDur,
			})
			bookingCount[earliestFreeLease.RoomNumber]++
			continue
		}

		roomNumber := heap.Pop(&availableMeetingRooms).(int)
		bookingCount[roomNumber]++

		heap.Push(&meetingRoomLeases, &MeetingRoomLease{
			RoomNumber: roomNumber,
			End:        meeting[1],
		})

	}

	var (
		maxID    = -1
		maxCount = math.MinInt
	)

	for id, c := range bookingCount {
		if c > maxCount {
			maxID = id
			maxCount = c
		}
		if c == maxCount && id < maxID {
			maxID = id
		}
	}
	return maxID
}
