import "slices"

func findClosestElements(arr []int, k int, x int) []int {
	i, _ := slices.BinarySearch(arr, x)
	var (
		leftArr  []int
		rightArr []int
	)

	l, r := i-1, i

	for ; l >= 0 && r < len(arr) && k > 0; k-- {
		if absDiff(x, arr[l]) < absDiff(x, arr[r]) {
			leftArr = append(leftArr, arr[l])
			l--
			continue
		}

		if absDiff(x, arr[l]) > absDiff(x, arr[r]) {
			rightArr = append(rightArr, arr[r])
			r++
			continue
		}

		if absDiff(x, arr[l]) == absDiff(x, arr[r]) {
			if arr[l] < arr[r] {
				leftArr = append(leftArr, arr[l])
				l--
			} else {
				rightArr = append(rightArr, arr[r])
				r++
			}
		}

	}

	for k > 0 && l >= 0 {
		leftArr = append(leftArr, arr[l])
		l--
		k--
	}

	for k > 0 && r < len(arr) {
		rightArr = append(rightArr, arr[r])
		r++
		k--
	}

	slices.Reverse(leftArr)
	return append(leftArr, rightArr...)
}

func absDiff(a, b int) int {
	if a >= b {
		return a - b
	}
	return b - a
}
