package LeetCodeHot100

func searchRange(nums []int, target int) []int {
	res := make([]int, 2)
	res = []int{-1, -1}
	start, end := 0, len(nums)
	for start < end {
		i := (start + end) / 2
		med := nums[i]
		if med == target {
			p1, p2 := i, i
			for p1 >= 0 && nums[p1] == target {
				res[0] = p1
				p1--
			}
			for p2 < len(nums) && nums[p2] == target {
				res[1] = p2
				p2++
			}
			return res
		} else if med < target {
			start = i + 1
		} else {
			end = i
		}
	}
	return res
}
