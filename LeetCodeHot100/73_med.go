package LeetCodeHot100

func setZeroes(matrix [][]int) {
	zeroes := make([]bool, len(matrix[0]))
	for i := 0; i < len(matrix); i++ {
		flag := 1
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 0 {
				zeroes[j] = true
				for a := i - 1; a >= 0; a-- {
					matrix[a][j] = 0
				}
				for b := j - 1; b >= 0; b-- {
					matrix[i][b] = 0
				}
				flag = 0
			}
			if zeroes[j] {
				matrix[i][j] = matrix[i][j] * flag * 0
			} else {
				matrix[i][j] = matrix[i][j] * flag
			}
		}
	}
}

/*
【题解】
一开始想用递归，下右递归改，上左迭代改，结果发现栈的后进先出特点，使得问题没有解决，且耗时长。
上述是在考虑写一个列的标记数组，以及每行的特殊标记，使用了m个存储空间，是不符合题目需求的。
看过官方题解后，突然意识到，根据我的想法，完全可以拿第一行和第一列作为标记数组，最后再考虑他们需不需要置零就可以了
（还是需要标志标记首行首列的状态）
*/

func setZeroes1(matrix [][]int) {
	var flagI, flagJ bool
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 0 {
				if i == 0 {
					flagI = true
				}
				if j == 0 {
					flagJ = true
				}
				for a := i - 1; j > 0 && a >= 0; a-- {
					matrix[a][j] = 0
				}
				for b := j - 1; i > 0 && b >= 0; b-- {
					matrix[i][b] = 0
				}
			}
			if i > 0 && j > 0 && (matrix[i][0] == 0 || matrix[0][j] == 0) {
				matrix[i][j] = 0
			}
		}
	}
	for i := 0; flagJ && i < len(matrix); i++ {
		matrix[i][0] = 0
	}
	for j := 0; flagI && j < len(matrix[0]); j++ {
		matrix[0][j] = 0
	}
}
