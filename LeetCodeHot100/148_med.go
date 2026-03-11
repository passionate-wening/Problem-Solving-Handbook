package LeetCodeHot100

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	newTail := head
	newHead := &ListNode{Val: 0, Next: newTail}
	head = head.Next
	for head != nil {
		suc := head.Next
		p := newHead
		for p != newTail {
			if head.Val < p.Next.Val {
				break
			}
			p = p.Next
		}
		head.Next = p.Next
		p.Next = head
		if p == newTail {
			newTail = p.Next
		}
		head = suc
	}
	newTail.Next = nil
	return newHead.Next
}

/*
【题解】
上述方法会超时...对于有序数组，其实属于O(n^2)了
//先找尾，新头、新尾指向这个尾
//遍历，暂存当前cur(head)、后继suc
//cur与新头到新尾链表检查小也生头，大则生尾
该方法有很多重复检查，官方给出两种：自顶向下归并；自底向上归并
只有自底向上归并的空间复杂度是O(1)，自顶向下是借用了栈空间。先写自顶向下的：
- 快慢指针找中点，记录首尾，断尾
- 二分到最后，单节点必然有序
- 两个有序链表，两两归并（同21题有序链表的合并）
*/

func sortList1(head *ListNode) *ListNode {
	return sortTwoList(head, nil)
}

func sortTwoList(head, tail *ListNode) *ListNode {
	if head == nil {
		return head
	}
	if head.Next == tail { //断尾才好合并
		head.Next = nil
		return head
	}
	var slow, fast *ListNode
	slow = head.Next
	fast = head.Next.Next
	for fast != tail && fast.Next != tail {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return mergeList(sortTwoList(head, slow), sortTwoList(slow, tail))
}

func mergeList(head1, head2 *ListNode) *ListNode {
	preNode := &ListNode{}
	p := preNode
	for head1 != nil && head2 != nil {
		if head1.Val <= head2.Val {
			p.Next = head1
			head1 = head1.Next
		} else {
			p.Next = head2
			head2 = head2.Next
		}
		p = p.Next
	}
	if head1 != nil {
		p.Next = head1
	} else if head2 != nil {
		p.Next = head2
	}
	return preNode.Next
}

/*
自底向上就是 两两归并->四四归并->八八归并...
一定要先算长度吗？我没算！我做对了！
*/

func sortList2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	flag := true
	subLen := 1
	preNode := &ListNode{Next: head}
	for flag {
		loopCount := 0
		pre := preNode
		cur := preNode.Next
		for cur != nil {
			head1 := cur
			for i := 1; i < subLen && cur.Next != nil; i++ {
				cur = cur.Next
			}
			head2 := cur.Next
			cur.Next = nil
			cur = head2
			for i := 1; i < subLen && cur != nil && cur.Next != nil; i++ {
				cur = cur.Next
			}
			var next *ListNode
			if cur != nil {
				next = cur.Next
				cur.Next = nil
			}
			pre.Next = mergeList(head1, head2)
			for pre.Next != nil {
				pre = pre.Next
			}
			cur = next
			loopCount++
		}
		if loopCount == 1 { //说明只合并了一次
			flag = false
		}
		subLen <<= 1
	}
	return preNode.Next
}
