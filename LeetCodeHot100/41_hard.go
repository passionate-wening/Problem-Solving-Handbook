package LeetCodeHot100

func firstMissingPositive(nums []int) int {
	for i, v := range nums {
		if v != i+1 {
			nums[i] = 0
			for v > 0 && v <= len(nums) && v != nums[v-1] {
				t := nums[v-1]
				nums[v-1] = v
				v = t
			}
		}
	}
	for i, v := range nums {
		if v == 0 {
			return i + 1
		}
	}
	return len(nums) + 1
}

/*
【题解】
最小正数就直接把数组当成从1开始的序列区间对比呗
判断i位数字，如果当前数字v不为i+1，则置0。v不在区间范围或重复，则直接丢弃；v在区间范围内且非重复，就覆盖v-1位数字，循环至被覆盖位被丢弃
再次循环取第一个为0的对应正数
好离谱的官解。。。。我觉得我的上述想法更好。
*/
