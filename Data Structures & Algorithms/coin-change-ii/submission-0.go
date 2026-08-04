func change(amount int, coins []int) int {
	sort.Ints(coins)
	mem := make([][]int,5001)
	for i := range mem {
		mem[i] = make([]int,len(coins))
		for j := range mem[i]{
			mem[i][j] = -1
		}
	}
    return changeWithMemo(amount,0, coins,mem)
}

func changeWithMemo(amount int,i int, coins []int, mem [][]int) int {
    if amount < 0 || i >= len(coins){
		return 0
	}

	if amount == 0 {
		return 1
	}

	if ways := mem[amount][i]; ways >= 0 {
		return ways
	}

	var totalWays int

	if c := coins[i]; amount >= c {
		totalWays += changeWithMemo(amount-c,i,coins,mem)
		totalWays += changeWithMemo(amount,i+1,coins,mem)
	}

	mem[amount][i] = totalWays

	return totalWays
}