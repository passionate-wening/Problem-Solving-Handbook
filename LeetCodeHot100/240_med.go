package LeetCodeHot100

func searchMatrix2(matrix [][]int, target int) bool {
	if matrix[0][0] == target {
		return true
	} else if matrix[0][0] > target {
		return false
	}
	row, col := len(matrix), len(matrix[0])
	i, j, ie, je := 0, 0, row-1, col-1
	for i < ie || j < je {
		flagI, flagJ := true, true
		for flagI || flagJ {
			l, r := matrix[ie][j], matrix[i][je]
			if l == target || r == target {
				return true
			} else {
				if l > target {
					ie--
				} else {
					flagI = false
				}
				if r > target {
					je--
				} else {
					flagJ = false
				}
			}
		}
		if i != ie {
			i++
		}
		if j != je {
			j++
		}
	}
	if matrix[i][j] == target {
		return true
	} else {
		return false
	}
}

/*
【题解】
找小矩阵：
- 从左下和右上向左上出发，找两个临近小值，为新矩阵的左下和右上；
- 从左下和右上向右下出发，找两个临近大值，为新矩阵的左下和右上
- 循环，有等值即输出，至同行/同列，对比中间内容
我上面图方便，始终向左上出发，找新角，正确，但效率差
官方题解，暴力做法跳过、每一行二分法跳过，题解三不错，z字形搜索
我关注了两个角，其实关注一个角就行了。从右上开始，大于目标值，j--；小于目标值，i++
*/

func searchMatrix2_3(matrix [][]int, target int) bool {
	i, j := 0, len(matrix[0])-1
	for i < len(matrix) && j >= 0 { //不要忘了等号
		if matrix[i][j] == target {
			return true
		} else if matrix[i][j] > target {
			j--
		} else if matrix[i][j] < target {
			i++
		}
	}
	return false
}
