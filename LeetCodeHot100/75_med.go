package LeetCodeHot100

func sortColors(nums []int) {
	p1, p2, i := -1, len(nums), 0
	for i != p2 {
		v := nums[i]
		if v == 0 {
			p1++
			nums[p1], nums[i] = nums[i], nums[p1]
		} else if v == 2 {
			p2--
			nums[p2], nums[i] = nums[i], nums[p2]
		} else if v == 1 {
			i++
		}
		if p1 == i {
			i++
		}
	}
}
