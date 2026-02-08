package LeetCodeHot100

func rotate(nums []int, k int) {

}

/*
【题解】
凯撒密码！
要求三种方式：
1）维护一个数组，可以完全复制，也可以只复制k个(注意边界)；
2）模运算依次替换，只携带temp，会死循环，得有个bool数组;
没想出来O(1)空间的，太困了，没有办法思考，和官方题解也不一致 ，后面再想
*/

func rotate1(nums []int, k int) {
	k = k % len(nums)
	bak := make([]int, len(nums))
	for i := 0; i < len(nums)-k; i++ {
		bak[i] = nums[i]
	}
	for i := len(nums) - k; i < len(nums); i++ {
		nums[i-len(nums)+k] = nums[i]
	}
	for i := k; i < len(nums); i++ {
		nums[i] = bak[i-k]
	}
}

func rotate2(nums []int, k int) {
	flag := make([]bool, len(nums))
	cur, next, temp := 0, k%len(nums), nums[0]
	for i := 0; i < len(nums); i++ {
		t := nums[next]
		nums[next] = temp
		flag[next] = true
		temp = t
		cur = next
		next = (cur + k) % len(nums)
		for flag[next] && i+1 < len(nums) {
			cur++
			next = (cur + k) % len(nums)
			temp = nums[cur]
		}
	}
}
