package LeetCodeHot100

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		if slow == fast {
			return true
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return false
}

/*
【题解】
考虑套圈情况，如果有环，龟兔赛跑一定能重新回合。
*/
/*
官方另一方法 ———— 哈希表，思路很简单，需要占用内存，不好，不记录了。
*/
