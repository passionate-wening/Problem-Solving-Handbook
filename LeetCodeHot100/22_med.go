package LeetCodeHot100

func generateParenthesis(n int) []string {
	res := []string{"()"}
	for i := 1; i < n; i++ {
		m := make(map[string]bool)
		for _, cur := range res {
			for j, s := range cur {
				if s == '(' {
					m[cur[:j]+"()"+cur[j:]] = true
					m[cur[:j+1]+"()"+cur[j+1:]] = true
				}
			}
		}
		res = []string{}
		for k := range m {
			res = append(res, k)
		}
	}
	return res
}

/*
【题解】
考虑递增序列，但要去重。也就是遍历上一个序列的所有左括号，把完整括号放在左括号前面、后面.
（）
（（））		（）（）
（）（（））	（（）（））	（（（）））	（）（）（）	（（））（）	（）（）（）	（）（（））
没必要dfs吧。。。我这个思路不是更灵活吗，，，
*/
