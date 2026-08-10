
type MemoKey [2]int

func findTargetSumWays(nums []int, target int) int {
	memo := make(map[MemoKey]int)

	return findTargetSumWaysFrom(nums, memo, 0, target)
}

func findTargetSumWaysFrom(nums []int, memo map[MemoKey]int, i int, target int) int {
	if i == len(nums) && target == 0 {
		return 1
	}

	if i >= len(nums) {
		return 0
	}

	key := MemoKey{i, target}

	if val, ok := memo[key]; ok {
		return val
	}

	memo[key] = findTargetSumWaysFrom(nums, memo, i+1, target-nums[i]) + findTargetSumWaysFrom(nums, memo, i+1, target+nums[i])
	return memo[key]
}
