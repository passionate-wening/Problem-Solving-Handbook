package LeetCodeHot100

func canJump(nums []int) bool {
	visited := make([]bool, len(nums))
	visited[0] = true
	for i, n := range nums {
		for j := 1; visited[i] && i+j < len(nums) && j <= n; j++ {
			visited[i+j] = true
		}
	}
	return visited[len(nums)-1]
}

/*
【题解】
上述做法是初次做法，思路闭塞，按照常规动态规划思路写的，维护一个visited，时间O(n)空间O(n)，很慢；
官解和网友的思路更好，我们不需要知道每一步的情况，只需要维护最远距离就可以了，见下↓
此做法最坏情况时间O(n)空间O(1)，未达终点情况时间O(k)空间O(1)；其实可以更好一点，多写一行只要计算到超过终点的距离就直接成功（我没写）。
*/

func canJump1(nums []int) bool {
	right := 0
	for i := 0; i <= right && i < len(nums); i++ {
		right = max(right, i+nums[i])
	}
	return right >= len(nums)-1
}
