package LeetCodeHot100

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func diameterOfBinaryTree(root *TreeNode) int {
	_, _, diameter := maxLength(root, 0)
	return diameter - 1 //由边数表示，diameter算的节点数
}

func maxLength(root *TreeNode, diameter int) (int, int, int) {
	if root == nil {
		return 0, 0, diameter
	}
	l1, r1, d1 := maxLength(root.Left, diameter)
	l2, r2, d2 := maxLength(root.Right, diameter)
	maxLeft := max(l1, r1)
	maxRight := max(l2, r2)
	d := maxLeft + maxRight + 1
	return maxLeft + 1, maxRight + 1, max(d, d1, d2)
}
