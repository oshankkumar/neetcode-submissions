
func maxProfit(prices []int) int {
	minSeen := math.MaxInt
	var maxProfit int
	for _, p := range prices {
		if p < minSeen {
			minSeen = p
			continue
		}
		maxProfit = max(maxProfit, p-minSeen)
	}
	return maxProfit
}
