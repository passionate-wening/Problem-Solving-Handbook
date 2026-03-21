package LeetCodeHot100

func minPathSum(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	visited := make([]int, m*n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			index := i*n + j
			temp := grid[i][j]
			visited[index] = temp
			if i-1 >= 0 {
				indexLog := (i-1)*n + j
				visited[index] = temp + visited[indexLog]
			}
			if j-1 >= 0 {
				indexLog := i*n + j - 1
				if visited[index] == temp {
					visited[index] = temp + visited[indexLog]
				} else {
					visited[index] = min(visited[index], temp+visited[indexLog])
				}
			}
		}
	}
	return visited[m*n-1]
}

/*
【题解】
动规到达每一个位置的最小值，所有数字大于0，来源只能是上面或左边，不然就不是最小。
官解是二维db，把首行首列都填充，再算剩下的；
我的性能更好，但要稍微复杂一点，我用一维表示二维，需要计算正确，同时没有预填写，需要过程中填写并比较，这里容易想错，要小心。
*/
