package LeetCodeHot100

func canPartition(nums []int) bool {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	if sum%2 != 0 {
		return false
	}
	sum = sum / 2
	x := 0
	var judge func(i int) bool
	judge = func(i int) bool {
		if i == len(nums) {
			return false
		}
		if x == sum {
			return true
		}
		x += nums[i]
		if judge(i + 1) {
			return true
		} else {
			x -= nums[i]
			if judge(i + 1) {
				return true
			} else {
				return false
			}
		}
	}
	return judge(0)
}

/*
【题解】
背包问题，要不要装满（这个东西装还是不装，贪心算法）。。。一直超时，非要我用dp
以后要重做，并没有非常理解。
*/

func canPartition1(nums []int) bool {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	if sum%2 != 0 {
		return false
	}
	sum = sum / 2
	dp := make([]bool, sum+1)
	dp[0] = true
	for _, n := range nums {
		for v := sum; v >= n; v-- {
			dp[v] = dp[v] || dp[v-n]
		}
	}
	return dp[sum]
}
