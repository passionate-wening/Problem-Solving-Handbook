package LeetCodeHot100

import (
	"sort"
)

func findKthLargest(nums []int, k int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})
	return nums[k-1]
}

/*
【题解】
维护一个k大的数组（我最快想到的内容了，但时间复杂度应该不是O(n)），得，我又往复杂了想，直接sort得了，其实是不合格的
之后要研究下快排和堆排
*/
