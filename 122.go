import "fmt"

func maxProfit(prices []int) int {
	min := prices[0]
	res := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] < prices[i-1] {
			fmt.Println("i", i)
			fmt.Println("res", res)
			fmt.Println("min", min)
			// 小的时候，卖
			res = max(res, res+(prices[i-1]-min))
			min = prices[i]
		} else {
			if i == len(prices)-1 {
				fmt.Println("in", i)
				fmt.Println("resn", res)
				fmt.Println("minn", min)
				res = max(res, res+(prices[i]-min))
			}
			continue
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}