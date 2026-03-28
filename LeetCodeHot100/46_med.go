package LeetCodeHot100

func permute(nums []int) [][]int {
	var res [][]int
	var temp []int
	visited := make([]bool, len(nums))
	var dfs func()
	dfs = func() {
		if len(temp) == len(nums) {
			res = append(res, append([]int{}, temp...))
			return
		}
		for i, n := range nums {
			if !visited[i] {
				temp = append(temp, n)
				visited[i] = true
				dfs()
				temp = temp[:len(temp)-1]
				visited[i] = false
			}
		}
	}
	dfs()
	return res
}

/*
【题解】
一直没写这个题，最开始思路不好总是写不全，后来熟练了dfs，所以准备写dfs了。需要注意去掉已经字符，写一个标记数组。
简单。注意好边界，什么时候加什么时候减什么时候改状态是什么时候是结束，就很快解决。
*/
