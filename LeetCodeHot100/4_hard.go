package LeetCodeHot100

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) == 0 {
		return findMed(nums2)
	} else if len(nums2) == 0 {
		return findMed(nums1)
	}
	nums := make([]int, len(nums1)+len(nums2))
	i, i1, i2 := 0, 0, 0
	for ; i < len(nums); i++ {
		if i1 >= len(nums1) {
			nums[i] = nums2[i2]
			i2++
		} else if i2 >= len(nums2) {
			nums[i] = nums1[i1]
			i1++
		} else if nums1[i1] <= nums2[i2] {
			nums[i] = nums1[i1]
			i1++
		} else {
			nums[i] = nums2[i2]
			i2++
		}
	}
	return findMed(nums)
}

func findMed(nums []int) float64 {
	if len(nums)%2 == 0 {
		return (float64(nums[len(nums)/2]) + float64(nums[len(nums)/2-1])) / 2.0
	}
	return float64(nums[len(nums)/2])
}

/*
这题要重做！投机取巧，不符合题目要求
*/
