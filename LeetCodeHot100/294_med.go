package LeetCodeHot100

import "strconv"

func decodeString(s string) string {
	var st []string
	start, n := 0, len(s)
	for i := 0; i < n; i++ {
		v := s[i : i+1]
		_, e := strconv.Atoi(v)
		if e == nil { //将数字前的字母与[前的数字压栈
			if start != i {
				st = append(st, s[start:i])
			}
			start = i
			for v != "[" {
				i++
				v = s[i : i+1]
			}
			st = append(st, s[start:i]) //只存数字
			start = i + 1
		} else if v == "]" {
			st = append(st, s[start:i]) //	先压栈，再找数字
			end := findNum(st)
			st = st[:end]
			start = i + 1
		}
	}
	res := ""
	for _, ss := range st {
		res += ss
	}
	if start < n {
		res += s[start:n]
	}
	return res
}

func findNum(st []string) int {
	end := len(st)
	for i := len(st) - 1; i >= 0; i-- {
		num, e := strconv.Atoi(st[i])
		if e == nil { //是数字，当前字母还未压栈
			r, cur := "", ""
			for j := i + 1; j < end; j++ {
				cur += st[j]
			}
			for j := 0; j < num; j++ {
				r += cur
			}
			st[i] = r
			return i + 1
		}
	}
	return -1
}

/*
【题解】
我跳过括号，直接压栈数字和字母，遇见右括号，直接找数字，拼好重复字母后覆盖数字，最后把所有字母连接起来。
有几个易错点：1、遇见左右括号后都要改下一段字符串的的起始位置；2、重复字母拼接直接正向顺序拼接；3、字符串首位可能为纯字母，不要遗漏。
和官解思路差不多，有一个点可以考虑，除了调用Atoi函数，还可以通过 cur >= '0' && cur <= '9' 判断数字
（我个人认为字母不用判断，只要不是括号和数字，按照规则，就只可能是字母）
不想看了，日后再好好复习下
*/
