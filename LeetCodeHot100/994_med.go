package LeetCodeHot100

func orangesRotting(grid [][]int) int {
	//找烂橘子位置、橘子总数
	nums, total := findLocation(grid)
	//多次腐蚀
	visited := make([][]int, len(grid))
	for i := range visited {
		visited[i] = make([]int, len(grid[0]))
	}
	for _, val := range nums {
		temp := make([][]int, len(grid)) //必须深拷贝
		for i := range visited {
			temp[i] = make([]int, len(grid[0]))
			copy(temp[i], grid[i])
		}
		visited = bfsRotting(temp, val/len(grid[0]), val%len(grid[0]), visited)
	}
	//遍历visited，1）有记录的总数小于橘子总数，-1；2）找时间最大值
	maxN, count := 0, 0
	for _, line := range visited {
		for _, val := range line {
			if val != 0 {
				count++
				if val > maxN {
					maxN = val
				}
			}
		}
	}
	if count < total {
		return -1
	} else {
		return maxN
	}
}

func findLocation(grid [][]int) ([]int, int) {
	var res []int
	count := 0
	for i, gLine := range grid {
		for j, val := range gLine {
			if val == 2 {
				res = append(res, i*len(grid[0])+j)
			} else if val == 1 {
				count++
			}
		}
	}
	return res, count
}

func bfsRotting(grid [][]int, i, j int, visited [][]int) [][]int {
	m := len(grid)
	n := len(grid[0])
	queue := [][]int{{i*n + j, 0}}
	for qi := 0; qi < len(queue); qi++ {
		cur := queue[qi]
		ci := cur[0] / n
		cj := cur[0] % n
		cc := cur[1]
		if grid[ci][cj] != 1 && (ci != i || cj != j) {
			continue
		}
		if visited[ci][cj] == 0 {
			visited[ci][cj] = cc
		} else {
			visited[ci][cj] = min(visited[ci][cj], cc)
		}
		grid[ci][cj] = 2
		if ci-1 >= 0 && grid[ci-1][cj] == 1 {
			next := (ci-1)*n + cj
			queue = append(queue, []int{next, cc + 1})
		}
		if cj-1 >= 0 && grid[ci][cj-1] == 1 {
			next := ci*n + cj - 1
			queue = append(queue, []int{next, cc + 1})
		}
		if ci+1 < m && grid[ci+1][cj] == 1 {
			next := (ci+1)*n + cj
			queue = append(queue, []int{next, cc + 1})
		}
		if cj+1 < n && grid[ci][cj+1] == 1 {
			next := ci*n + cj + 1
			queue = append(queue, []int{next, cc + 1})
		}
	}
	return visited
}

/*
【题解】
我的想法是，先遍历计算好橘子的个数、坏橘子的位置。
维护一个二维表格，从坏橘子出发，腐蚀周围的橘子，如果该橘子历史腐蚀时间更短，则终止腐蚀；否则覆盖
维护一个最大腐蚀时间
需要bfs而不是dfs；数组要深拷贝！不然浅拷贝的数组牵扯变动
虽然做对了，但思路过于复杂，且空间(2mn+logn)visited\temp\queue，时间O(mn)
上述做法是从单源出发，做bfs，考虑初始多个橘子腐烂，采用多源bfs算法，即初始烂橘子视为同一层（官解）其实只需要一个queue
下述借鉴了网友的做法，更简洁，性能更好一些
*/

func orangesRotting1(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	var queue []int
	fresh, maxTime := 0, 0
	for i, line := range grid {
		for j, val := range line {
			if val == 2 {
				queue = append(queue, i*n+j)
			} else if val == 1 {
				fresh++
			}
		}
	}
	bfs := func(i, j int) {
		if i >= 0 && j >= 0 && i < m && j < n && grid[i][j] == 1 {
			queue = append(queue, i*n+j)
			grid[i][j] = 2
			fresh--
		}
	}
	for len(queue) > 0 && fresh > 0 {
		cur := queue
		queue = []int{} //每次抓一层，就可以按层计数了
		for _, v := range cur {
			vi := v / n
			vj := v % n
			bfs(vi+1, vj)
			bfs(vi-1, vj)
			bfs(vi, vj+1)
			bfs(vi, vj-1)
		}
		maxTime++
	}
	if fresh > 0 {
		maxTime = -1
	}
	return maxTime
}
