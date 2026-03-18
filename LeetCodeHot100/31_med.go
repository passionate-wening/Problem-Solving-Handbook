package LeetCodeHot100

func nextPermutation(nums []int) {
	for i := len(nums) - 1; i > 0; i-- {
		if nums[i] > nums[i-1] {
			sortBubble(nums, i)
			j := i
			for ; i < len(nums); j++ {
				if nums[j] > nums[i-1] {
					break
				}
			}
			nums[j], nums[i-1] = nums[i-1], nums[j]
			return
		}
	}
	sortBubble(nums, 0)
	return
}

func sortBubble(nums []int, i int) {
	for ; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
}

/*
【题解】
从右向左，右>左，逆转，后面冒泡从小到大
若无右>左，则直接从小到大冒泡

1、冒泡排序不要写错
2、参考网友题解，找到“右>左”不要直接逆转，先排序“右”及“右”后面的，再逆转第一个比“左”大的。

复杂度较高，有时间看下题解
*/
