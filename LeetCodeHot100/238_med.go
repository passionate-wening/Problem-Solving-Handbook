package LeetCodeHot100

func productExceptSelf(nums []int) []int {
	res := make([]int, len(nums))
	product, index, count := 1, -1, 0
	for i, n := range nums {
		if n == 0 {
			count++
			index = i
			if count > 1 {
				break
			}
		} else {
			product *= n
		}
	}
	if count == 1 {
		res[index] = product
	} else if count == 0 {
		for i, n := range nums {
			res[i] = product / n
		}
	}
	return res
}

/*
【题解】
算出总乘积，除掉当前就可以了。
需要区分的是0的情况，如果只有一个0，记录当前位置算其他乘积，其余为0；如果超过1个0，则结果全0
我没懂官方题解和网友在卷什么，这个题很复杂吗，后面再研究下好了。
*/
