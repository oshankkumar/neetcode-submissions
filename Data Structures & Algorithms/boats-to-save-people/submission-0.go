
func numRescueBoats(people []int, limit int) int {
	sort.Ints(people)

	var count int
	left, right := 0, len(people)-1

	for left < right {
		sum := people[left] + people[right]
		if sum <= limit {
			left++
			right--
		} else {
			right--
		}
		count++
	}

	if left == right {
		count++
	}
	return count
}
