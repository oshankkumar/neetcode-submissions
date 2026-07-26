
import (
	"slices"
)

func findClosestElements(arr []int, k int, x int) []int {
	i := sort.Search(len(arr), func(i int) bool {
		return arr[i] >= x
	})
	if i == len(arr) {
		return arr[len(arr)-k:]
	}

	if i == 0 {
		return arr[0:k]
	}

	l, r := i-1, i

	var (
		leftArr  []int
		rightArr []int
	)

	for k > 0 && l >= 0 && r < len(arr) {
		if absDiff(x, arr[l]) < absDiff(x, arr[r]) {
			leftArr = append(leftArr, arr[l])
			l--
		} else if absDiff(x, arr[l]) > absDiff(x, arr[r]) {
			rightArr = append(rightArr, arr[r])
			r++
		} else if arr[l] < arr[r] {
			leftArr = append(leftArr, arr[l])
			l--
		} else {
			rightArr = append(rightArr, arr[r])
			r++
		}
		k--
	}

	for l >= 0 && k > 0 {
		leftArr = append(leftArr, arr[l])
		l--
		k--
	}

	for r < len(arr) && k > 0 {
		rightArr = append(rightArr, arr[r])
		r++
		k--
	}

	slices.Reverse(leftArr)

	return append(leftArr, rightArr...)
}

func absDiff(x, y int) int {
	if x > y {
		return x - y
	}
	return y - x
}
