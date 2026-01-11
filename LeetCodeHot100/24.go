package LeetCodeHot100

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return nil
	} else if head.Next == nil { //单数情况
		return head
	} else {
		p1, p2 := head, head.Next
		p1.Next = swapPairs(p2.Next)
		p2.Next = p1
		return p2
	}
}

/*
【题解】
考虑拆成子任务，递归。始终认为后续链为正确的，只交换前两个节点，返回第二个节点。
挺简单的，不算中等难度。
官方题解也是递归或迭代，其实迭代递归的另一种表现形式，之所以空间复杂度要小，是因为递归用了栈空间。也可以在下面写写迭代。
一个很关键的思想，head之前写头节点！！很多题通过这个方式都会更方便一点，一定要记住!
*/

func swapPairs1(head *ListNode) *ListNode {
	preHead := &ListNode{Val: 0, Next: head}
	temp := preHead //站在前面的视角交换后面两个
	for temp.Next != nil && temp.Next.Next != nil {
		p1 := temp.Next
		p2 := temp.Next.Next
		temp.Next = p2
		p1.Next = p2.Next
		p2.Next = p1
		temp = p1
	}
	return preHead.Next
}
