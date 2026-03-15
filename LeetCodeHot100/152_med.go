package LeetCodeHot100

func maxProduct(nums []int) int {
	res := nums[0]
	maxV, minV := res, res
	for i := 1; i < len(nums); i++ {
		v := nums[i]
		temp := minV
		minV = min(min(minV*v, v), maxV*v)
		maxV = max(max(maxV*v, v), temp*v)
		res = max(res, maxV)
	}
	return res
}

/*
【题解】
好吧，还是看了网友做法：维护最大、最小值的动态规划
*/
