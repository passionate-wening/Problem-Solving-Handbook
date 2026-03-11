package LeetCodeHot100

func subsets(nums []int) [][]int {
	for i := 0; i <= len(nums); i++ {
		subTemp = make([]int, i)
		sub(nums, 0)
	}
	newRes := make([][]int, len(subRes))
	copy(newRes, subRes)
	subRes = [][]int{}
	return newRes
}

var subRes [][]int
var subTemp []int

func sub(nums []int, i int) {
	if i == len(subTemp) {
		r := make([]int, len(subTemp))
		copy(r, subTemp)
		subRes = append(subRes, r)
	} else {
		for j, n := range nums {
			subTemp[i] = n
			sub(nums[j+1:], i+1)
		}
	}
}

/*
【题解】
以上是蠢办法。
记录官解的两个妙计，思路一致，每一个元素都只有两种可能性：要么选择、要么不选择。
题解一：元素两种可能性化为0或1，按照原序列顺序组合，答案的选择方式就是 0到2^n-1的二进制序列（左移右移的算术方法很值得学习）
题解二：从无出发每个元素的两种可能性化为二叉树，答案就是dfs遍历
网友的思路很有趣！
网友：线性动态规划，从空序列开始，每轮向当前已有序列中新增一个元素，形成新的一组序列，并到已有的序列组中，指数增长。
*/

func subsets1(nums []int) [][]int {
	var res [][]int
	n := len(nums)
	for mark := 0; mark < 1<<n; mark++ {
		var temp []int
		for i, v := range nums {
			if mark>>i&1 == 1 {
				temp = append(temp, v)
			}
		}
		res = append(res, temp)
	}
	return res
}

func subsets2(nums []int) (res [][]int) {
	var temp []int
	var dfs func(int)
	dfs = func(cur int) {
		if cur == len(nums) {
			res = append(res, append([]int(nil), temp...))
			return
		}
		temp = append(temp, nums[cur])
		dfs(cur + 1)
		temp = temp[:len(temp)-1]
		dfs(cur + 1)
	}
	dfs(0)
	return
}

func subsets3(nums []int) [][]int {
	var res [][]int
	res = append(res, []int(nil))
	for i := 0; i < len(nums); i++ {
		var temp [][]int
		for _, r := range res {
			temp = append(temp, append(append([]int(nil), r...), nums[i]))
		}
		res = append(res, temp...)
	}
	return res
}
