package LeetCodeHot100

func trap(height []int) int {
	leftH := make([]int, len(height))
	rightH := make([]int, len(height))
	leftH[0] = height[0]
	rightH[len(height)-1] = height[len(height)-1]
	for i := 1; i < len(height); i++ {
		leftH[i] = max(leftH[i-1], height[i])
		rightH[len(height)-1-i] = max(rightH[len(height)-i], height[len(height)-1-i])
	}
	sum := 0
	for i := 0; i < len(height); i++ {
		sum += min(leftH[i], rightH[i]) - height[i]
	}
	return sum
}

func min(m, n int) int {
	if m < n {
		return m
	} else {
		return n
	}
}
func max(m, n int) int {
	if m > n {
		return m
	} else {
		return n
	}
}

/*
【题解】
看了一眼思路，醍醐灌顶！
要有分解的思想，我们最终要算雨水含量，其实是每一个点有可能的水量之和，而每一个点的水量=min(左边最高柱子，右边最高柱子)-当前柱子高度！
这思路挺好的，就是有点占内存，官方题解暂时没读，有精神的时候再学习。
*/
