package LeetCodeHot100

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseList_(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	p1 := head.Next.Next
	p2 := head.Next
	p3 := head
	p3.Next = nil
	for p1 != nil {
		p2.Next = p3
		p3 = p2
		p2 = p1
		p1 = p1.Next
	}
	p2.Next = p3
	return p2
}

/*
【题解】
其实这个思路不难，无非就是遍历的同时把指针反过来就可以了，但我写的比官方答案复杂很多，为什么呢？因为我把第三个指针当成遍历的主体了。
如果把遍历这个任务交给中间的指针，提前创建前面的空指针，过程中暂存未来的后继结点指针，整体写出来要简单得多。
（注意：边界条件是当前点走到空值，暂存点只在进行时有效）
*/

func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	curr := head
	for curr != nil {
		suc := curr.Next
		curr.Next = pre
		pre = curr
		curr = suc
	}
	return pre
}

/*
官方给出的另一个题解是递归，可以借此复习下子任务思想。
子任务是：curr.Next.Next = curr ; curr.Next = nil （及时断链，避免成环）
边界条件是：（末）curr.Next == nil ; （空链）curr == nil
递归的目的是倒着改指针，并记录末尾节点所串接的新链
n1 > n2 > ... > nk > nk+1 < ... < nm
				↓		↓
				nk < nk+1
但其实这相当于遍历了两遍链，不如迭代。
*/

func reverseList2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	suc := reverseList2(head.Next)
	head.Next.Next = head
	head.Next = nil
	return suc
}
