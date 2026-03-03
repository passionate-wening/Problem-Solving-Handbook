package LeetCodeHot100

func findMin(nums []int) int {
	start, end := 0, len(nums)-1
	if start == end || nums[start] < nums[end] {
		return nums[start]
	}
	if end == 1 {
		return nums[end]
	}
	for start <= end {
		med := (start + end) / 2
		pre := med - 1 //已经排除0了，所以不可能小于0
		if nums[pre] > nums[med] {
			return nums[med]
		} else if nums[med] < nums[end] {
			end = med
		} else {
			start = med + 1
		}

	}
	return -1
}
