func maxProfit(prices []int, fee int) int {
	n := len(prices)
	dp := make([][2]int, n)
	dp[0][1] = -prices[0]
	for i := 1; i < n; i++ {
		dp[i][0] = max(dp[i-1][0],dp[i-1][1]+prices[i]-fee)
		dp[i][1] = max(dp[i-1][0]-prices[i], dp[i-1][1])
	}
	return dp[n-1][0]
}

func max(a,b int) int{
	if a > b {
		return a
	}
	return b
}

func maxProfit(prices []int, fee int) int {
    n := len(prices)
    sell, buy := 0, -prices[0]
    for i := 1; i < n; i++ {
        sell = max(sell, buy+prices[i]-fee)
        buy = max(buy, sell-prices[i])
    }
    return sell
}


