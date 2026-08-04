func coinChange(coins []int, amount int) int {
	memo := make([]int,10000+1)
	for i := range memo{
		memo[i] = math.MaxInt
	}
    return coinChangeMemo(coins,amount,memo )
}

func coinChangeMemo(coins []int, amount int, memo []int) int{
	if amount <= 0 {
		return amount
	}

	if memo[amount] != math.MaxInt {
		return memo[amount]
	}

	coinCount := math.MaxInt

	for _,coin := range coins {
		count := coinChangeMemo(coins,amount-coin,memo) 
		if count < 0 {
			continue
		}
		coinCount = min(coinCount,count+1)
	}

	if coinCount == math.MaxInt {
		memo[amount] = -1
	} else {
		memo[amount] =  coinCount
	}

	return memo[amount]
}
