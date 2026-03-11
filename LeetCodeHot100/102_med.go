package LeetCodeHot100

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	var results [][]int
	var queue []*TreeNode
	queue = append(queue, root)
	for len(queue) > 0 {
		var temp []int
		var newQueue []*TreeNode
		for _, node := range queue {
			temp = append(temp, node.Val)
			if node.Left != nil {
				newQueue = append(newQueue, node.Left)
			}
			if node.Right != nil {
				newQueue = append(newQueue, node.Right)
			}
		}
		results = append(results, temp)
		queue = newQueue
	}
	return results
}

/*
简单的层序遍历，写出来和官方一致，记住这叫：广度优先搜索
*/
