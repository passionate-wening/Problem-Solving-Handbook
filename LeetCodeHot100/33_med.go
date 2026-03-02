package LeetCodeHot100

func search(nums []int, target int) int {
	start, end := 0, len(nums)-1
	for start <= end {
		if nums[start] < nums[end] {
			return findTarget(nums, start, end, target)
		} else {
			med := (start + end) / 2
			if nums[med] == target {
				return med
			} else if nums[start] == target {
				return start
			} else if nums[end] == target {
				return end
			} else if nums[start] < nums[med] {
				if nums[start] < target && target < nums[med] {
					return findTarget(nums, start+1, med-1, target)
				} else {
					start = med + 1
				}
			} else { // nums[med]<nums[end]
				if nums[med] < target && target < nums[end] {
					return findTarget(nums, med+1, end-1, target)
				} else {
					end = med - 1
				}
			}
		}
	}
	return -1
}

// 普通二分查找
func findTarget(nums []int, start, end, target int) int {
	for start <= end {
		med := (start + end) / 2
		v := nums[med]
		if v == target {
			return med
		} else if v > target {
			end = med - 1
		} else {
			start = med + 1
		}
	}
	return -1
}

/*
【题解】
旋转有序数组，局部单增，首尾为原始首尾/相邻数字
原始首尾(nums[start]<nums[end])：
	普通二分查找
相邻数字(nums[start]>nums[end])：判断med位置，
	若 nums[start]<nums[med] && nums[start]<target<nums[med]，则普通二分
	若   nums[med]<nums[end] && nums[med]<target<nums[end]，则普通二分
	其他情况循环上述判断
看了官解，一遍二分就可以
*/

func search1(nums []int, target int) int {
	start, end := 0, len(nums)-1
	for start <= end {
		med := (start + end) / 2
		if nums[med] == target {
			return med
		}
		if nums[0] <= nums[med] {
			if nums[0] <= target && target < nums[med] {
				end = med - 1
			} else {
				start = med + 1
			}
		} else {
			if nums[med] < target && target <= nums[len(nums)-1] {
				start = med + 1
			} else {
				end = med - 1
			}
		}
	}
	return -1
}
