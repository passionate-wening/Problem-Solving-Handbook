package LeetCodeHot100

func combinationSum(candidates []int, target int) [][]int {
	var res [][]int
	var temp []int
	var cal func(i, demand int)
	cal = func(i, demand int) {
		if demand == 0 {
			res = append(res, append([]int{}, temp...))
			return
		}
		if demand < 0 || i >= len(candidates) {
			return
		}
		temp = append(temp, candidates[i])
		cal(i, demand-candidates[i])
		temp = temp[:len(temp)-1]
		cal(i+1, demand)
	}
	cal(0, target)
	return res
}

/*
【题解】
能想到贪心算法，先拿大的，从大向小往多了拿，有效值放备选数组：确定，存答案；遍历完，备选数组中删小的，再顺序往下遍历（想复杂了）
dfs写不对，还是问了ai，是思路没想好，要考虑这个问题是要不要把当前元素重复放入口袋的问题，然后贪心增减
很经典的题解，不能再忘了
*/
