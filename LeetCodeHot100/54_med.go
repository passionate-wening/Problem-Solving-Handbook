package LeetCodeHot100

import "fmt"

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return nil
	}
	var res []int
	flag := 0
	m := make(map[string]bool)
	m[fmt.Sprintf("i%d", len(matrix))] = true
	m[fmt.Sprintf("j%d", len(matrix[0]))] = true
	m["i-1"] = true
	m["j-1"] = true
	count, i, j := 0, 0, -1
	for flag < 2 { //碰壁两次就结束
		if count%4 == 0 {
			j++
			if m[fmt.Sprintf("j%d", j)] {
				m[fmt.Sprintf("i%d", i)] = true
				count++
				j--
				flag++
				continue
			}
		} else if count%4 == 1 {
			i++
			if m[fmt.Sprintf("i%d", i)] {
				m[fmt.Sprintf("j%d", j)] = true
				count++
				i--
				flag++
				continue
			}
		} else if count%4 == 2 {
			j--
			if m[fmt.Sprintf("j%d", j)] {
				m[fmt.Sprintf("i%d", i)] = true
				count++
				j++
				flag++
				continue
			}
		} else if count%4 == 3 {
			i--
			if m[fmt.Sprintf("i%d", i)] {
				m[fmt.Sprintf("j%d", j)] = true
				count++
				i++
				flag++
				continue
			}
		}
		res = append(res, matrix[i][j])
		flag = 0
	}
	return res
}

/*
【题解】
考虑碰壁旋转90°：右-增加j，下-增加i，左-减小j，上-减小i
模运算就可以轮转了
map记录遍历过的i/j
时间复杂度O(mn)，空间复杂度O(m+n)
等有时间看看官方题解
*/
