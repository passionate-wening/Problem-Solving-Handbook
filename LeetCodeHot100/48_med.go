package LeetCodeHot100

func rotate_48(matrix [][]int) {
	n := len(matrix)
	for i := 0; i < n/2; i++ {
		for j := i; j < n-1-i; j++ {
			rotate_mod(matrix, i, j)
		}
	}
}

func rotate_mod(matrix [][]int, i, j int) {
	n := len(matrix) - 1
	temp := matrix[i][j]
	matrix[i][j] = matrix[n-j][i]
	matrix[n-j][i] = matrix[n-i][n-j]
	matrix[n-i][n-j] = matrix[j][n-i]
	matrix[j][n-i] = temp
}

/*
【题解】
单点旋转，转四个点，比如：
1,1   1,n-1   n-1,n-1 n-1,1
0,0   0,n     n,n     n,0
0,1   0+1,n   n,n-1   n-1,0
0,2   0+2,n   n,n-2   n-2,0
0,n-1 0+n-1,n n,1     1,0
只要把这些位置分清，就没有问题，这次算是一把通了，唯一疏漏的是外层j遍历，忘了n-1
和官方题解第一种方法基本一模一样，他是按角遍历的，我是按一半的行数遍历的，旋转是一个意思，别转多了就行
题解三很有趣啊，是数学方法，看起来更聪明：先水平翻转再垂直翻转，记录一下：
*/

func rotate_48_3(matrix [][]int) {
	n := len(matrix)
	for i := 0; i < n/2; i++ {
		for j := 0; j < n; j++ {
			temp := matrix[i][j]
			matrix[i][j] = matrix[n-i-1][j]
			matrix[n-i-1][j] = temp
		}
	}
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			temp := matrix[i][j]
			matrix[i][j] = matrix[j][i]
			matrix[j][i] = temp
		}
	}
}

/*
答案和我的区别就是它利用了go的特性，可以多元素同时替换。（直接行替换其实不符合题目要求，空间复杂度要O(1)）
for i := 0; i < n/2; i++ {
	matrix[i], matrix[n-1-i] = matrix[n-1-i], matrix[i]
}
该操作是‌直接交换两个切片的元数据‌，而不是复制底层数组。
Go 的赋值操作是‌按值交换‌，但因为切片是引用类型，所以交换的是切片的元数据（指针、长度、容量），而不是底层数组的内容。
因此，这种交换是‌非常高效的‌，不需要额外的内存分配或元素复制。
我认为本质是go底层替我封装了内层元素遍历复制。
*/
