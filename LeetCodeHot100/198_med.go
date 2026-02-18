package LeetCodeHot100

func rob(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	} else if n == 2 {
		return max(nums[0], nums[1])
	}
	visited := make([]int, n)
	visited[0] = nums[0]
	visited[1] = nums[1]
	visited[2] = max(nums[0], nums[2])
	for i := 3; i < n; i++ {
		visited[i] = max(visited[i-2], visited[i-3]) + nums[i]
	}
	return max(visited[n-1], visited[n-2])
}
