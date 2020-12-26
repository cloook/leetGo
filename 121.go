func maxProfit(prices []int) int {
	minPrice := math.MaxInt32
	maxMoney := 0
	for i := 0; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		} else {
			maxMoney = max(maxMoney, prices[i]-minPrice)
		}
	}
	return maxMoney
}