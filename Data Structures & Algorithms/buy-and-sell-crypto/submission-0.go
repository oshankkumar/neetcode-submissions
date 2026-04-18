func maxProfit(prices []int) int {
    var (
		maxProfit       int
		currentMinPrice int = 1<<31 - 1
	)

	for _, p := range prices {
		if p < currentMinPrice {
			currentMinPrice = p
			continue
		}

		maxProfit = max(maxProfit, p-currentMinPrice)
	}

	return maxProfit
}
