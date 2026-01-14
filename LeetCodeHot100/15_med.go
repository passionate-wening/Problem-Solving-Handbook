package LeetCodeHot100

import "sort"

func threeSum(nums []int) [][]int {
	var result [][]int
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	completed := make(map[int]bool)
	count := 0
	for i := 0; i < len(nums) && nums[i] <= 0; i++ {
		if nums[i] == 0 {
			count++
			continue
		}
		target := nums[i]
		if _, ok := completed[target]; !ok {
			re := checkTwoSum(nums[i+1:], target)
			result = append(result, re...)
			completed[target] = true
		}
	}
	if count >= 3 {
		result = append(result, []int{0, 0, 0})
	}
	return result
}

func checkTwoSum(nums []int, target int) (result [][]int) {
	m := make(map[int]bool)
	completed := make(map[int]bool)
	for _, n := range nums {
		if _, ok := completed[n]; ok {
			continue
		}
		r := 0 - target - n
		if _, ok := m[r]; ok {
			result = append(result, []int{target, n, r})
			delete(m, r)
			completed[n] = true
			completed[r] = true
		} else {
			m[n] = true
		}
	}
	return result
}

/*
【题解】
先考虑了排序+正数负数双指针遍历，结果总是控制不好；
后来考虑以前写过的两数之和，考虑拆解题目，遍历一个数作为target，在剩下的数组中找两数之和，结果出现了大量重复数据，
为了解决这一问题，我又增加了完成控制、重复数据跳过机制，越写越复杂了，于是就有了上面这个试错多次得到的题解。
不是好思想，目前在力扣报错最多的一道题。
官方题解是更加繁琐愚蠢的层层遍历，个人认为不是一个好题，或者说没有找到好的解决办法，太困了，下次再研究...
*/
