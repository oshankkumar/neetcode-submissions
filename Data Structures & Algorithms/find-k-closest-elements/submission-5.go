import "slices"

func findClosestElements(arr []int, k int, x int) []int {
	i, _ := slices.BinarySearch(arr, x)

	l, r := i-1, i

	for r-l-1 < k {
		switch {
		case l < 0:
			r++
		case r >= len(arr):
			l--
		case distance(x, arr[l]) <= distance(x, arr[r]):
			l--
		default:
			r++
		}
	}

	return arr[l+1 : r]
}

func distance(a, b int) int {
	if a >= b {
		return a - b
	}
	return b - a
}