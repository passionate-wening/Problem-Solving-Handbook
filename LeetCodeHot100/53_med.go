package LeetCodeHot100

func maxSubArray(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		dp[i] = max(nums[i], nums[i]+dp[i-1])
	}
	res := -1 << 31
	for _, s := range dp {
		if s > res {
			res = s
		}
	}
	return res
}

/*
【题解】
状态不好，看了思路，动态规划，我的死穴...
dp存以当前位置（i）结尾的最大统计和。
官方题解原地改数组。。这倒是不浪费O(n)空间了。。。
分治法没看
看到一个聪明的网友：前缀和————维护一个最小的前缀和，始终计算当前前缀和，同时统计当前前缀和与最小前缀和的差值，就是最大连续数组
下次遇到“连续和”，记得考虑前缀和
*/
