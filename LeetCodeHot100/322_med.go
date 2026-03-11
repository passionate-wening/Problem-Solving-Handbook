package LeetCodeHot100

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	dp[0] = 0
	for i := 0; i <= amount; i++ {
		dp[i] = amount + 1
		for _, c := range coins {
			if i-c > 0 && dp[i-c] >= 0 {
				dp[i] = min(dp[i], dp[i-c]+1)
			}
		}
		if dp[i] == amount+1 {
			dp[i] = -1
		}
	}
	return dp[amount]
}
