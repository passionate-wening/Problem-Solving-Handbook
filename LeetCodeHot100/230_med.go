package LeetCodeHot100

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthSmallest(root *TreeNode, k int) int {
	res := -1
	var preSort func(root *TreeNode)
	preSort = func(root *TreeNode) {
		if root == nil {
			return
		}
		preSort(root.Left)
		k--
		if k == 0 {
			res = root.Val
			return
		}
		preSort(root.Right)
	}
	preSort(root)
	return res
}

/*
【题解】
前序遍历，用迭代法，存k个栈时，就停止遍历并输出
官方题解过于复杂了，为了学习的话可以日后看看，其实有点过度设计
*/
