package LeetCodeHot100

func lengthOfLIS(nums []int) int {
	res := 0
	visited := make([]int, len(nums))
	for i, n := range nums {
		visited[i] = 1
		j := i
		for ; j >= 0; j-- {
			if nums[j] < n {
				visited[i] = max(visited[i], visited[j]+1)
			}
		}
		res = max(res, visited[i])
	}
	return res
}

/* 1,3,6,7,9,4,10,5,6
【题解】
考虑visited数组记录当前位置为重点的最长子序列。向前遍历，接在最长的序列后面。
时间复杂度O(n^2)，空间复杂度O(n)，效率不好。解法与官解一一致。
没有好好看第二种，有时间好好看一下
*/
