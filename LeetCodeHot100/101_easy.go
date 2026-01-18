package LeetCodeHot100

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
	if root.Left == nil {
		if root.Right == nil {
			return true
		} else {
			return false
		}
	}
	return isSame(root.Left, root.Right)
}

func isSame(t1, t2 *TreeNode) bool {
	if t1 == nil || t2 == nil {
		if t2 != nil || t1 != nil {
			return false
		} else {
			return true
		}
	}
	if t1.Val != t2.Val {
		return false
	}
	isLeftSame := isSame(t1.Left, t2.Right)
	isRightSame := isSame(t1.Right, t2.Left)
	return isLeftSame && isRightSame
}

/*
【题解】
写复杂了，官方题解更简单。
*/

func isSymmetric1(root *TreeNode) bool {
	return checkSymmetric1(root.Left, root.Right)
}

func checkSymmetric1(t1, t2 *TreeNode) bool {
	if t1 == nil && t2 == nil {
		return true
	}
	if t1 == nil || t2 == nil {
		return false
	}
	return t1.Val == t2.Val && checkSymmetric1(t1.Left, t2.Right) && checkSymmetric1(t1.Right, t2.Left)
}

/*
下次试一下写迭代，不想写了
*/
