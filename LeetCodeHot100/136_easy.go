package LeetCodeHot100

func singleNumber(nums []int) int {
	res := 0
	for _, n := range nums {
		res = res ^ n
	}
	return res
}

/*
【题解】
直接想到异或
*/
