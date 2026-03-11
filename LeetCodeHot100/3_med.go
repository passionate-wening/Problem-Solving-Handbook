package LeetCodeHot100

func lengthOfLongestSubstring(s string) int {
	m := make(map[string]int)
	left, right := 0, 0
	maxLength := 0
	for right < len(s) {
		c := s[right : right+1]
		if _, ok := m[c]; ok {
			for c != s[left:left+1] {
				delete(m, s[left:left+1])
				left++
			}
			delete(m, s[left:left+1])
			left++
		}
		m[c] = right
		right++
		if right-left > maxLength {
			maxLength = right - left
		}
	}
	return maxLength
}

/*
【题解】
已知属于滑动窗口组，我就按照我理解的双指针做了。
字符串截断是左闭右开，右边的边界一直走，map查重，有重复则左边删除map至重复元素的下一个。
官方题解和我写的差不多，都不太好，删除map是耗时的；我截取字符串，官方是转byte，差不多，可能写法清晰一点
我都让哈希表存位置了，但没用上，网友的建议不错，哈希表不删除，左边界定位到重复元素的下一位，检查重复时判断是否在滑动窗口内。
测试性能差不多，没删map顶多就占点微不足道的内存，节省点微不足道的map删除时间。
*/

func lengthOfLongestSubstring1(s string) int {
	m := make(map[byte]int)
	left, right := 0, 0
	maxLength := 0
	for right < len(s) {
		c := s[right]
		if index, ok := m[c]; ok && index >= left {
			left = index + 1
		}
		m[c] = right
		right++
		if right-left > maxLength {
			maxLength = right - left
		}
	}
	return maxLength
}
