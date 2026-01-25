package LeetCodeHot100

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var res []int
	var stack []*TreeNode
	stack = append(stack, root)
	for len(stack) != 0 {
		res = append(res, stack[len(stack)-1].Val)
		var temp []*TreeNode
		for _, node := range stack {
			if node != nil && node.Left != nil {
				temp = append(temp, node.Left)
			}
			if node != nil && node.Right != nil {
				temp = append(temp, node.Right)
			}
		}
		stack = temp
	}
	return res
}

/*
【题解】
层序遍历，输出最右就好了
*/
