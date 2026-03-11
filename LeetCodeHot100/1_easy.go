package LeetCodeHot100

func twoSum(nums []int, target int) []int {
	var res []int
	i := 0
	for ; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				res = append(res, i)
				res = append(res, j)
				return res
			}
		}
	}
	return res
}

/*
【题解】
我没想到更好的方式，无非是空间换时间。
官方题解是建立哈希表快速定位。
如果只是少量内存占用换取时间的话，是有意义的；但是如果大量内存占用换时间就不值了。
目前来看，这个换取还是有点意义。性能至上。
*/

func twoSum1(nums []int, target int) []int {
	hashSet := make(map[int]int)
	for i, val := range nums {
		n := target - val
		if j, ok := hashSet[n]; ok {
			return []int{i, j}
		} else {
			hashSet[val] = i
		}
	}
	return []int{}
}
