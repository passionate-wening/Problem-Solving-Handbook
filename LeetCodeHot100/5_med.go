package LeetCodeHot100

func longestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	} else if len(s) == 1 {
		return s
	}
	var check func(s string, i, j int) string
	check = func(s string, i, j int) string {
		for s[i] == s[j] {
			i--
			j++
			if i < 0 || j >= len(s) {
				break
			}
		}
		return s[i+1 : j]
	}
	var res string
	for i := 0; i < len(s)-1; i++ {
		temp1 := check(s, i, i)
		temp2 := check(s, i, i+1)
		if len(res) < len(temp1) {
			res = temp1
		}
		if len(res) < len(temp2) {
			res = temp2
		}
	}
	return res
}

/*
【题解】
太晚了，直接看题解用其中最简单的办法吧：中心扩展判断
*/
