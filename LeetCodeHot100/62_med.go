package LeetCodeHot100

func uniquePaths(m int, n int) int {
	visited := make([][]int, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]int, n)
		visited[i][0] = 1
	}
	for j := 0; j < n; j++ {
		visited[0][j] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			visited[i][j] = visited[i-1][j] + visited[i][j-1]
		}
	}
	return visited[m-1][n-1]
}

/*
【题解】
两个方向动态规划
可恶，递归就超时，就是得写visited
不够冷静，写得不够好，再练
*/
