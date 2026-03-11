package LeetCodeHot100

func partitionLabels(s string) []int {
	m := make(map[string]int)
	for j := 0; j < len(s); j++ {
		m[s[j:j+1]] = j
	}
	start, end, i, n := 0, 0, 0, len(s)
	var res []int
	for i < n {
		end = max(end, m[s[i:i+1]])
		if i == end {
			res = append(res, end-start+1)
			start, end = i+1, i+1
		}
		i++
	}
	return res
}

/*
【题解】
和跳跃游戏Ⅱ45题是一样的逻辑。我一开始写的map（都被当作示例了）。题解用的[26]int，总是忘记这种表达字母的方法，记录下这种。
*/

func partitionLabels1(s string) []int {
	var m [26]int
	for i, ss := range s {
		m[ss-'a'] = i
	}
	var res []int
	start, end := 0, 0
	for i, ss := range s {
		end = max(end, m[ss-'a'])
		if end == i {
			res = append(res, end-start+1)
			start, end = i+1, i+1
		}
		i++
	}
	return res
}
