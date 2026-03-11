package LeetCodeHot100

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func pathSum(root *TreeNode, targetSum int) int {
	return preTreeRecursion(root, 0, targetSum, map[int]int{0: 1})
}

func preTreeRecursion(node *TreeNode, preVal, target int, record map[int]int) int {
	if node == nil {
		return 0
	}
	cur := 0
	node.Val += preVal
	if r, ok := record[node.Val-target]; ok {
		cur += r
	}
	record[node.Val]++
	res := preTreeRecursion(node.Left, node.Val, target, record) +
		preTreeRecursion(node.Right, node.Val, target, record) + cur
	record[node.Val]--
	if record[node.Val] == 0 {
		delete(record, node.Val)
	}
	return res
}

/*
【题解】
联想到之前的路径和，考虑前缀树和+map，难点是如何建立树状的前缀树和。
写对了！中间犯错点在于：前缀和取差值要注意减数和被减数关系；注意正好等于目标值的情况，初始化条件携带0
官方题解一是暴力，不提倡，题解二与我一致
*/
