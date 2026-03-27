package LeetCodeHot100

func partition(s string) [][]string {
	dp := make([][]bool, len(s))
	for i := range dp {
		dp[i] = make([]bool, len(s))
	}
	for i := len(s) - 1; i >= 0; i-- {
		for j := i; j < len(s); j++ {
			dp[i][j] = s[i] == s[j]
			if dp[i][j] && j > i+2 {
				dp[i][j] = dp[i][j] && dp[i+1][j-1]
			}
		}
	}

	var res [][]string
	var temp []string
	var dfs func(i int)
	dfs = func(i int) {
		if i == len(s) {
			res = append(res, append([]string{}, temp...))
			return
		}
		for j := i; j < len(s); j++ {
			if dp[i][j] {
				temp = append(temp, s[i:j+1])
				dfs(j + 1)
				temp = temp[:len(temp)-1]
			}
		}
	}
	dfs(0)
	return res
}

/*
【题解】
历史没做对，本次参考题解。动态规划思路很妙，i、j表示判断回文串的边界，历史是回文串就只需要判断新的就可以了。
首先动规初始化所有的回文串组（这里网友的思路更好，所以我记录了网友的：二维数组对角线下侧左下三角区域i > j，根据dp的定义，这是没有意义的区域；针对右上三角，从下往上倒着递推）
然后dfs做分割判断

以上，要多复习！
*/
