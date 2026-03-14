package LeetCodeHot100

func wordBreak(s string, wordDict []string) bool {
	m := make(map[string]bool)
	for _, word := range wordDict {
		m[word] = true
	}
	return checkWord(s, 0, 1, m)
}

func checkWord(s string, start, end int, dict map[string]bool) bool {
	if start >= len(s) {
		return true
	}
	for end <= len(s) {
		_, ok := dict[s[start:end]]
		if ok {
			if checkWord(s, end, end+1, dict) {
				return true
			}
		}
		end++
	}
	return false
}

/*
【题解】
截断查map，超时。
看了题解，判断前缀树是否合法
*/

func wordBreak1(s string, wordDict []string) bool {
	m := make(map[string]bool)
	for _, word := range wordDict {
		m[word] = true
	}
	visited := make([]bool, len(s)+1)
	visited[0] = true
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if visited[j] && m[s[j:i]] {
				visited[i] = true
				break
			}
		}
	}
	return visited[len(s)]
}
