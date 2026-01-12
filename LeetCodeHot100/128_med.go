package LeetCodeHot100

import "sort"

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	start, end, maxLength := 0, 1, 1 // [start,end) maxLength = end-start
	for ; end < len(nums); end++ {
		if nums[end-1] == nums[end] {
			start++
		} else if nums[end-1]+1 != nums[end] {
			if maxLength < (end - start) {
				maxLength = end - start
			}
			start = end
		}
	}
	if maxLength < (end - start) {
		maxLength = end - start
	}
	return maxLength
}

/*
【题解】
双指针，连续数组不会重叠，不断更新最大长度就可以了；要先排序；还要去重...直接让前面的指针进一位，反正目的是统计；
15ms	9.5MB
官方题解个人感觉也一般，使用哈希表，倒是省去了排序去重，思路更简单，记得遍历map，不然会超时。
48ms	12MB
整体效果，不如我写的，使用map既占内存，还耗时
*/

func longestConsecutive1(nums []int) int {
	m := make(map[int]bool)
	maxLength := 0
	for _, n := range nums {
		m[n] = true
	}
	for n := range m {
		if m[n-1] {
			continue
		}
		count := 0
		for m[n] {
			count++
			n++
		}
		if maxLength < count {
			maxLength = count
		}
	}
	return maxLength
}
