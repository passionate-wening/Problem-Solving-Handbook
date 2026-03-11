package LeetCodeHot100

func numIslands(grid [][]byte) int {
	count := 0
	for i, gline := range grid {
		for j, val := range gline {
			if val == '1' {
				count++
				dfsIsland(grid, i, j)
			}
		}
	}
	return count
}

func dfsIsland(grid [][]byte, i, j int) {
	if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[0]) || grid[i][j] == '0' {
		return
	}
	grid[i][j] = '0'
	dfsIsland(grid, i+1, j)
	dfsIsland(grid, i-1, j)
	dfsIsland(grid, i, j+1)
	dfsIsland(grid, i, j-1)
}

/*
【题解】
看了东哥思路，好聪明，计数后直接把岛淹了！DFS四叉树，无需维护visited数组！
*/
