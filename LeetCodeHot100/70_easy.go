package LeetCodeHot100

func climbStairs(n int) int {
	if n == 1 {
		return 1
	} else if n == 2 {
		return 2
	}
	visited := make([]int, n)
	visited[0] = 1
	visited[1] = 2
	for i := 2; i < n; i++ {
		visited[i] = visited[i-2] + visited[i-1]
	}
	return visited[n-1]
}

/*
【题解】
动态规划经典爬楼梯
有时间研究下题解。
*/
