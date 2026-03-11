package LeetCodeHot100

func generate(numRows int) [][]int {
	if numRows == 1 {
		return [][]int{{1}}
	} else if numRows == 2 {
		return [][]int{{1}, {1, 1}}
	}
	res := [][]int{}
	res = append(res, []int{1})
	res = append(res, []int{1, 1})
	for i := 2; i < numRows; i++ {
		r := make([]int, i+1)
		r[0] = 1
		r[i] = 1
		for j := 1; j < i; j++ {
			r[j] = res[i-1][j-1] + res[i-1][j]
		}
		res = append(res, r)
	}
	return res
}
