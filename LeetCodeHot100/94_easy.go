package LeetCodeHot100

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	var res []int
	res = append(res, inorderTraversal(root.Left)...)
	res = append(res, root.Val)
	res = append(res, inorderTraversal(root.Right)...)
	return res
}

/*
【题解】中序遍历：左->中->右
是的，递归很简单，要会用迭代写。
起初想复杂了，还自己写链表模拟栈，思路不对
参考官方的迭代写法，维护栈可以用数组的
*/

func inorderTraversal1(root *TreeNode) []int {
	var stack []*TreeNode
	var res []int
	for root != nil || len(stack) > 0 { //左侧清空，取每一个中间节点，再右移
		for root != nil {
			stack = append(stack, root) //入栈，只存中间节点
			root = root.Left
		}
		root = stack[len(stack)-1]   //取中间节点
		stack = stack[:len(stack)-1] //出栈
		res = append(res, root.Val)
		root = root.Right
	}
	return res
}
