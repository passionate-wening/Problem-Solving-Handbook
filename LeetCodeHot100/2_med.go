package LeetCodeHot100

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	p1 := l1
	p2 := l2
	n := 0
	sum := 0
	head := &ListNode{Val: 0}
	p := head
	for p1 != nil || p2 != nil {
		if p1 == nil {
			sum = p2.Val + n
			p2 = p2.Next
		} else if p2 == nil {
			sum = p1.Val + n
			p1 = p1.Next
		} else {
			sum = p1.Val + p2.Val + n
			p1 = p1.Next
			p2 = p2.Next
		}
		l := &ListNode{Val: sum % 10}
		p.Next = l
		p = p.Next
		n = sum / 10
	}
	if n != 0 {
		p.Next = &ListNode{Val: n}
	}
	return head.Next
}
