package LeetCodeHot100

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{Val: preorder[0]}
	med := 0
	for i, val := range inorder {
		if val == root.Val {
			med = i
			break
		}
	}
	root.Left = buildTree(preorder[1:len(inorder[:med])+1], inorder[:med])
	root.Right = buildTree(preorder[len(inorder[:med])+1:], inorder[med+1:])
	return root
}

/*
【题解】
先序遍历可以确定当前树的根节点，中序遍历可以根据根节点，确定当前的左右子树
思路是清楚的，就是递归，但脑子不清楚，所以去看官方了，以后要自己做出来
*/
