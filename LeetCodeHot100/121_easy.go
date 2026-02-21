package LeetCodeHot100

func maxProfit(prices []int) int {
	minPrice, profit := prices[0], 0
	for i := 1; i < len(prices); i++ {
		if minPrice >= prices[i] {
			minPrice = prices[i]
		} else {
			profit = max(profit, prices[i]-minPrice)
		}
	}
	return profit
}
