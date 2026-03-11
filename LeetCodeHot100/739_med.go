package LeetCodeHot100

func dailyTemperatures(temperatures []int) []int {
	n := len(temperatures)
	res := make([]int, n)
	res[n-1] = 0
	for i := 2; i <= n; i++ {
		cur := n - i
		temp := cur + 1
		for temperatures[cur] >= temperatures[temp] {
			if res[temp] == 0 {
				break
			}
			temp = temp + res[temp]
		}
		if temperatures[cur] < temperatures[temp] {
			res[cur] = temp - cur
		}
	}
	return res
}

/*
【题解】
可以倒着找更高的温度，一遍通过。虽然是一道栈练习，但我觉得没必要欸。
学习一下栈做法：
*/

func dailyTemperatures1(temperatures []int) []int {
	n := len(temperatures)
	res := make([]int, n)
	st := make([]int, 1)
	st[0] = 0
	for i := 1; i < n; i++ { //完全可以从0开始，不给st分配空间，下面的逻辑会自动压栈的
		top := len(st) - 1
		for top >= 0 && temperatures[st[top]] < temperatures[i] {
			v := st[top]
			st = st[:top]
			res[v] = i - v
			top--
		}
		st = append(st, i)
	}
	for _, s := range st { //这里不写也默认0
		res[s] = 0
	}
	return res
}
