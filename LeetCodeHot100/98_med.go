package LeetCodeHot100

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
	return isValidT(root.Left, math.MinInt, root.Val) && isValidT(root.Right, root.Val, math.MaxInt)
}

func isValidT(root *TreeNode, min, max int) bool {
	if root == nil {
		return true
	} else if root.Val <= min || root.Val >= max || (root.Left != nil && root.Left.Val >= root.Val) || (root.Right != nil && root.Right.Val <= root.Val) {
		return false
	}
	return isValidT(root.Left, min, root.Val) && isValidT(root.Right, root.Val, max)
}

/*
【题解】
严格！要知道什么是严格！
左边的所有子树都不能大于中间值，右边的所有子树都不能小于中间值
等于也不可以！！
递归要记得改界限
*/
