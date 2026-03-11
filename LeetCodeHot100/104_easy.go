package LeetCodeHot100

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

/*
【题解】
巨简单，go也有自己的模板max/min函数。但你要懂这种方法叫“深度优先搜索”。
官方给出方法二：广度优先搜索。是一种迭代写法，下次再看吧。
*/
