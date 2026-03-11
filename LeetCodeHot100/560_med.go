package LeetCodeHot100

// 这是一段错误的解答
func subarraySumX(nums []int, k int) int {
	flag, next := true, true //首轮扩大数组
	left, right, sum, count := 0, 0, 0, 0
	preDiff := abs(sum, k)
	for flag {
		diff := abs(sum, k)
		if sum == k && left != right { //无差距时做什么都增大差距，所以数据增大
			count++
		} else if diff > preDiff { //差距增大，改变变动方式
			next = !next
		}
		sum, left, right, flag = change(nums, left, right, sum, next)
		preDiff = diff
	}
	return count
}

func abs(a, b int) int {
	if a >= b {
		return a - b
	} else {
		return b - a
	}
}

func change(nums []int, left, right, sum int, flag bool) (int, int, int, bool) { // true 扩大 false 缩小 [left,right)
	if flag {
		if right >= len(nums) {
			return 0, 0, 0, false
		}
		return sum + nums[right], left, right + 1, true
	} else {
		if left >= len(nums) {
			return 0, 0, 0, false
		}
		return sum - nums[left], left + 1, right, true
	}
}

/*
【题解】
不是纯粹的滑动窗口，因为k有正有负；
考虑每一次扩大/缩小对结果的影响，维护一个结果与当前和的差距距离，若变动使差距增大，则转换变动方式，若变动使差距减小，则继续该变动
可恶，有负数存在，就不能用这种方式，写个[1,2,-3]就挂了
真的不想枚举吧....
官方题解二好聪明啊，我没有想到过“前缀和”。
考虑滑动窗口是着眼于窗口边缘，总数上下波动是很难去判断扩容还是缩容的；
前缀和是一种宏观思维，考虑我们要取的连续数组，实际上是某位置的完整长度-前面的某段完整长度，即[x,y] = [0,y]-[0,x]
这样只需要维护一个前缀和map就可以，key为“前缀和”，value为“出现次数”
*/

func subarraySum(nums []int, k int) int {
	m := make(map[int]int)
	m[0] = 1
	sum, count := 0, 0
	for _, n := range nums {
		sum += n
		if c, ok := m[sum-k]; ok {
			count += c
		}
		m[sum] += 1
	}
	return count
}

/*妙*/
