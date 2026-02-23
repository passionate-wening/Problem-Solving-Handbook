package LeetCodeHot100

func jump(nums []int) int {
	end, maxEdge, step := 0, 0, 0
	for i := 0; i < len(nums)-1; i++ {
		maxEdge = max(maxEdge, i+nums[i])
		if maxEdge >= len(nums)-1 {
			step++
			break
		}
		if i == end {
			end = maxEdge
			step++
		}
	}
	return step
}

/*
我又下意识写了visited数组。同55题，我们只需要维护最远距离，再多维护几个数字：步数、当前边界
所以不记录笨方法了，记录正向贪心。（没有必要算最后一位，算上反而会影响0步跳跃情况）
*/
