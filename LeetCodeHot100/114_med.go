package LeetCodeHot100

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	var stack []*TreeNode //后进先出
	stack = append(stack, root)
	p := &TreeNode{Val: 0}
	for len(stack) != 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		p.Right = node
		p = p.Right
		for p.Right != nil || p.Left != nil {
			if p.Left != nil {
				if p.Right != nil {
					stack = append(stack, p.Right)
				}
				p.Right = p.Left
				p.Left = nil
				p = p.Right
			} else if p.Right != nil {
				p = p.Right
			}
		}
	}
}

/*
【题解】
每轮：
	出栈，向右连接当前节点，循环：
		左边有节点，就存非空的右节点，向右连接左边，并清空左边，右移；
		左边无节点，但右边有节点，右移；
		左右都无节点，下一轮
后进先出，栈非空则继续轮次。
*/
