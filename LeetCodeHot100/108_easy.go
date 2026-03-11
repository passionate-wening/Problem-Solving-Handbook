package LeetCodeHot100

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedArrayToBST(nums []int) *TreeNode {
	return createTree(&TreeNode{}, nums)
}

func createTree(root *TreeNode, nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	med := len(nums) / 2
	root.Val = nums[med]
	root.Left = createTree(&TreeNode{}, nums[:med])
	root.Right = createTree(&TreeNode{}, nums[med+1:])
	return root
}

/*
【题解】
有序数组，直接默认取中点就得了。数组取值，左闭右开。
我上来就考虑了乱序数组怎么排，感觉都写成堆排序了，但我不应该出问题，需要找机会好好check下下面的代码。
*/
func sortedArrayToBST1(nums []int) *TreeNode {
	root := &TreeNode{Val: nums[0]}
	hl, hr := 0, 0
	for i := 1; i < len(nums); i++ {
		h := 0
		p := root
		flag := false
		for !flag {
			if nums[i] > p.Val {
				if p.Right == nil {
					flag = true
					p.Right = &TreeNode{Val: nums[i]}
				} else {
					p = p.Right
				}
			} else if nums[i] < p.Val {
				if p.Left == nil {
					flag = true
					p.Left = &TreeNode{Val: nums[i]}
				} else {
					p = p.Left
				}
			}
			h++
		}
		if nums[i] > root.Val {
			hr = max(hr, h)
		} else {
			hl = max(hl, h)
		}
		root, hr, hl = adjust(root, hr, hl)
	}
	return root
}

func adjust(root *TreeNode, hr, hl int) (*TreeNode, int, int) {
	if hl+2 == hr { //右边长
		newRoot := root.Right
		temp := newRoot.Left
		newRoot.Left = root
		root.Right = temp
		root = newRoot
		hl = heightCount(root.Left) + 1
		hr = heightCount(root.Right) + 1
	} else if hl == hr+2 { //左边长
		newRoot := root.Left
		temp := newRoot.Right
		newRoot.Right = root
		root.Left = temp
		root = newRoot
		hl = heightCount(root.Left) + 1
		hr = heightCount(root.Right) + 1
	}
	return root, hr, hl
}

func heightCount(root *TreeNode) (h int) {
	if root == nil {
		return 0
	}
	return max(heightCount(root.Left), heightCount(root.Right)) + 1
}
