package LeetCodeHot100

func searchMatrix(matrix [][]int, target int) bool {
	lines := len(matrix)
	columns := len(matrix[0])
	start, end := 0, lines
	for start+1 != end && start != end {
		i := (end + start) / 2
		val := matrix[i][0]
		if val == target {
			return true
		} else if val < target {
			start = i
		} else {
			end = i
		}
	}
	resI := 0
	if end == lines {
		resI = start
	} else if target == matrix[end][0] || target == matrix[start][0] {
		return true
	} else if target > matrix[end][0] {
		resI = end
	} else if target > matrix[start][0] {
		resI = start
	} else {
		return false
	}
	start, end = 0, columns
	for start+1 != end {
		j := (end + start) / 2
		val := matrix[resI][j]
		if val == target {
			return true
		} else if val < target {
			start = j
		} else {
			end = j
		}
	}
	if matrix[resI][start] == target || (end != columns && matrix[resI][end] == target) {
		return true
	} else {
		return false
	}
}

/*
【题解】
二分法找定位：先找行，后找列。
虽然写对了，但写的好复杂，官方题解一也是两次二分查，题解二是一次二分查，即视作一维升序数组，逻辑映射上去
*/

func searchMatrix1(matrix [][]int, target int) bool {
	lines := len(matrix)
	columns := len(matrix[0])
	total := lines * columns
	start, end := 0, total-1
	for start <= end {
		m := (start + end) / 2
		i, j := m/columns, m%columns
		val := matrix[i][j]
		if target == val {
			return true
		} else if target > val {
			start = m + 1
		} else {
			end = m - 1
		}
	}
	return false
}

/*代码阅读复杂度立竿见影的降低了*/
