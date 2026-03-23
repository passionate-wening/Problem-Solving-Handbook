package LeetCodeHot100

func longestCommonSubsequence(text1 string, text2 string) int {
	var temp [26]bool
	for _, n := range text1 {
		temp[n-'a'] = true
	}
	s := ""
	for _, n := range text2 {
		if temp[n-'a'] {
			s += string(n)
		}
	}
	res := 0
	var findIndex func(start int, target uint8) int
	findIndex = func(start int, target uint8) int {
		for i := start; i < len(text1); i++ {
			if text1[i] == target {
				return i
			}
		}
		return -1
	}
	var bag func(i int, t []int)
	bag = func(i int, t []int) {
		if i >= len(s) {
			if len(t) > res {
				res = len(t)
			}
			return
		}
		if len(t) == 0 {
			t = append(t, findIndex(0, s[i]))
			bag(i+1, t)
			t = t[:len(t)-1]
			bag(i+1, t)
		} else {
			in := findIndex(t[len(t)-1]+1, s[i])
			if in != -1 {
				t = append(t, in)
				bag(i+1, t)
				t = t[:len(t)-1]
				bag(i+1, t)
			} else {
				bag(i+1, t)
			}
		}
	}
	bag(0, []int{})
	return res
}

/*
【题解】
先统计一个字母范围，顺便记录首次出现位置吧，再拿另一个搜索，把不存在的字母剔除，重组，再背包，拿这个和不拿这个的长度。
看你写的复杂度就知道，超时，呵
还得是dp，还得是状态转移，官解的状态转移思路真的很好：记录两个前缀子串的最大公共长度。
if test1[i-1]==test2[j-1] => dp[i][j]=dp[i-1][j-1]+1
else => dp[i][j]=max(dp[i-1][j],dp[i][j-1])
*/

func longestCommonSubsequence1(text1 string, text2 string) int {
	m, n := len(text1), len(text2)
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if text1[i] == text2[j] {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				dp[i+1][j+1] = max(dp[i][j+1], dp[i+1][j])
			}
		}
	}
	return dp[m][n]
}

/*
我服了，边界边界边界，写dp就怕写错边界，一个是初始化有没有初始化全（是i<=m），一个是真正的遍历是遍历的两个text，所以是从0开始，而dp是要大一位数的，这些都看清楚看清楚看清楚！！
*/
