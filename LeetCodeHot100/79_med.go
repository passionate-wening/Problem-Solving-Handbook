package LeetCodeHot100

func exist(board [][]byte, word string) bool {
	towards := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	var check func(visited []bool, i, j, k int) bool
	check = func(visited []bool, i, j, k int) bool {
		index := i*len(board[0]) + j
		if k == len(word) {
			return true
		}
		if i < 0 || j < 0 || i >= len(board) || j >= len(board[0]) || visited[index] || board[i][j] != word[k] {
			return false
		}
		visited[index] = true
		for _, t := range towards {
			if check(visited, i+t[0], j+t[1], k+1) {
				return true
			}
		}
		visited[index] = false
		return false
	}
	for i, line := range board {
		for j, v := range line {
			if v == word[0] {
				if check(make([]bool, len(board)*len(line)), i, j, 0) {
					return true
				}
			}
		}
	}
	return false
}

/*
【题解】
先看看会不会超时：查到第一个元素匹配，就递归check，记录一个visited
不会超时，但要注意i、j、k边界、visited清理、二维转一维位置计算（是i*列数+j）！！！
经典回溯，我写的和题解一模一样！而且我写的比官解更好，我的visited用的一维，计算方法一定别写错！
*/
