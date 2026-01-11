package LeetCodeHot100

func moveZeroes(nums []int) {
	index1 := 0
	index2 := 0
	for ; index1 < len(nums); index1++ {
		if nums[index1] == 0 {
			continue
		}
		nums[index2] = nums[index1]
		index2++
	}
	for ; index2 < len(nums); index2++ {
		nums[index2] = 0
	}
}

/*
【题解】
典型双指针问题，要做到不debug快速通过。
官方题解是位置交换，我写的左移填零。
- 左移填零优势：赋值的操作比交换的操作少;
- 官方优势：零值多时，遍历次数少。
下面的答案比官方多一步两个指针位置判断，减少多余操作。
*/

func moveZeroes1(nums []int) {
	index1 := 0
	index2 := 0
	for ; index1 < len(nums); index1++ {
		if nums[index1] != 0 {
			if index1 != index2 {
				nums[index1], nums[index2] = nums[index2], nums[index1]
			}
			index2++
		}
	}
}
