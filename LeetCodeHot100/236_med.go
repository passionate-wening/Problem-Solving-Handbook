package LeetCodeHot100

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == p.Val || root.Val == q.Val {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left != nil && right != nil {
		return root
	}
	if left != nil {
		return left
	}
	return right
}

/*
【题解】
瞄了一眼思路。对于一个根节点，1）左右一边一个，返回结果；2）单一边右边有，返回节点；3）都没有，返回null
其实重要的不是左右节点是什么，重要的是根。对根来说，确定有无节点就可以了。
后面要去学习一下这一类题，东哥引据git的merge和rebase，需要了解深切一点。
官方题解的递归和该方法一模一样，非递归方法日后看看吧，听说有地方面试要多种算法（虽然没必要，但有机会看看吧）
*/
