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
func maxPathSum(root *TreeNode) int {
	res := math.MinInt64
	maxNodeSum(root, &res)
	return res
}

func maxNodeSum(root *TreeNode, maxSum *int) int {
	if root == nil {
		return 0
	}
	leftMax := maxNodeSum(root.Left, maxSum)
	rightMax := maxNodeSum(root.Right, maxSum)
	*maxSum = max(*maxSum, rightMax+root.Val)
	*maxSum = max(*maxSum, leftMax+root.Val)
	*maxSum = max(*maxSum, root.Val)
	sum := leftMax + rightMax + root.Val
	*maxSum = max(*maxSum, sum)
	return max(max(leftMax, rightMax)+root.Val, root.Val)
}

/*
【题解】
每一个节点为根的最大路径和=左边最长+右边最长+根
从叶子向上递归，算当前节点和，比对最大路径和后，向上爬并携带当前最长边+当前根，直到遍历完所有的节点
题目挺简单，但做错了很多遍才做对。因为考虑得不周到：
1）初始化应该是最小值，2）没有判断纯节点无边情况，3）没有判断单边情况，4）没有考虑中间段最大情况，递归最长边的时候，也要考虑纯根情况
*/
