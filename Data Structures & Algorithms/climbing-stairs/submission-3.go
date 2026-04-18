func climbStairs(n int) int {
	return climbStairsUtil(n, make(map[int]int))
}

func climbStairsUtil(n int, mem map[int]int) int {
	if n <= 2 {
		return n
	}

	if val, ok := mem[n]; ok {
		return val
	}


	mem[n] = climbStairsUtil(n-1, mem) + climbStairsUtil(n-2, mem) 

	return mem[n]
}
