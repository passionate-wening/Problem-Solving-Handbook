package LeetCodeHot100

import "strings"

func solveNQueens(n int) [][]string {
	var res [][]string
	var temp []string
	var dfs func(i int)
	dfs = func(i int) {
		if i == n-1 {
			res = append(res, append([]string{}, temp...))
			return
		}
		//下一行的选择，确定可行再dfs
		for j := 0; j < n; j++ {
			var isQ func(i, j, k int) bool
			isQ = func(i, j, k int) bool {
				if j-i+k >= 0 && j-i+k < n && temp[k][j-i+k] == 'Q' {
					return true
				} else if j+i-k >= 0 && j+i-k < n && temp[k][j+i-k] == 'Q' {
					return true
				} else {
					return false
				}
			}
			flag := true
			for k := i; k >= 0; k-- {
				if temp[k][j] == 'Q' || isQ(i+1, j, k) {
					flag = false
					break
				}
			}
			if flag {
				temp = append(temp, strings.Repeat(".", j)+"Q"+strings.Repeat(".", n-j-1))
				dfs(i + 1)
				temp = temp[:len(temp)-1]
			}
		}
	}
	dfs(-1)
	return res
}

/*
【题解】
既然N个放在N*N里，行、列、对角线不交叉，那么，每行至少得有一个才行，所以遍历完一行的位置就够了。
维护一个visited，每放下一个子，就把对应行、列斜线填实。
有个问题，取消这个选择，怎么清除visited记录，所以这不是一个好的回溯。
那就不维护了，每一个取值的时候都要去确定是否可用，但问题是，这将是非常坏的性能，但题目1<=n<=9，说明不太大，是不是可以这样做，只要探上面几行就行了
写对了！

怎么算的对角线：k按行计算
	首先看对角线 \ ，正对角线是k==i==j，
		那么i和j的差值就是偏移量，
		左下角(同i，j减小，i>j)使对角线左移，需要减差值，刚好j-i就是负数
		右上角(同i，j增大，i<j)使对角线右移，需要加上差值，刚好j-i就是正数
		算列，所以从（原j）k偏移，就是k+(j-i)=j-i+k
	然后看对角线 / ，负对角线是k==i==n-1-j,
		那么i和n-1-j的差值就是偏移量，
		左上角(同i，j减小，i<n-1-j）使对角线左移，需要减差值，刚好i-(n-1-j)就是负数
		右下角（同i，j增大，i>n-1-j）使对角线右移，需要加差值，刚好i-(n-1-j)就是正数
		算列，所以从（原j）n-1-k偏移，就是n-1-k+[i-(n-1-j)]=j+i-k


不太想看官解，下次吧，我这个算法也不难，就是思考得久一点。
*/
