package LeetCodeHot100

func findAnagrams(s string, p string) []int {
	anagrams := anagramsSign(p)
	left, right := 0, len(p)
	var results []int
	for right <= len(s) {
		word := s[left:right]
		if anagrams == anagramsSign(word) {
			results = append(results, left)
		}
		left++
		right++
	}
	return results
}

func anagramsSign(s string) string {
	var sign [26]byte
	for i := 0; i < len(s); i++ {
		sign[s[i]-'a']++
	}
	return string(sign[:])
}

/*
【题解】又是异位词，去翻了49题，已经忘记当时总结的好做法了。。。
对异位词签名。缺点是，每三个数字都要判断，如果异位词足够长，滑动窗口期间，中间大部分是不需要重新判断的，只需要判断左边少的和右边多的
官方题解的性能要好一些，确实难写一点。
先写性能普通的方法一：维护26位数组，窗口边界增加/减少，复制一份数组，做对应位的增加/减少，比较数组差异
*/

func findAnagrams1(s string, p string) []int {
	if len(s) < len(p) {
		return []int{}
	}
	var anagrams, newAnagrams [26]int
	var results []int
	for i, v := range p {
		anagrams[v-'a']++
		newAnagrams[s[i]-'a']++
	}
	if anagrams == newAnagrams {
		results = append(results, 0)
	}
	for left, v := range s[:len(s)-len(p)] { //开始移动，即循环内的真实区间是[left+1,left+len(p)]
		newAnagrams[v-'a']--
		right := left + len(p)
		newAnagrams[s[right]-'a']++
		if anagrams == newAnagrams {
			results = append(results, left+1)
		}
	}
	return results
}

/*
方法二是进一步优化，我们不用两个数组，只需要一个数组和差异数量就可以了，滑动窗口发生的变动会产生差异，有种汉明算法的感觉
*/

func findAnagrams2(s string, p string) []int {
	if len(s) < len(p) {
		return []int{}
	}
	var count [26]int
	var results []int
	for i, v := range p { //一定要让测试组-对照组，这样测试组后续增减才是对的。负数是少了；正数是多余
		count[s[i]-'a']++
		count[v-'a']--
	}
	diff := 0
	for _, c := range count {
		if c != 0 {
			diff++
		}
	}
	if diff == 0 {
		results = append(results, 0)
	}
	for left, v := range s[:len(s)-len(p)] {
		count[v-'a']--
		if count[v-'a'] == 0 { //差异消除
			diff--
		} else if count[v-'a'] == -1 { //差异增加，要写清楚，因为可能有重复数据影响结果
			diff++
		}
		right := left + len(p)
		count[s[right]-'a']++
		if count[s[right]-'a'] == 0 { //差异消除
			diff--
		} else if count[s[right]-'a'] == 1 { //差异增加，同上
			diff++
		}
		if diff == 0 {
			results = append(results, left+1)
		}
	}
	return results
}
