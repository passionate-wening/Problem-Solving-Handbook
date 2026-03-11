package LeetCodeHot100

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	p1 := headA
	for p1 != nil {
		p2 := headB
		for p2 != nil {
			if p1 == p2 { //内存一致是可以这样判断的
				return p1
			}
			p2 = p2.Next
		}
		p1 = p1.Next
	}
	return nil
}

/*
【题解】
O(m*n)是笨办法。
瞄了一眼题解，更好的O(m+n)方法是：两条链表相连，将长短不同的链表变成相同长度就可以了。
官方代码更漂亮（方法二），所以下面附上修改版本（没写边界条件是觉得没必要）
*/

func getIntersectionNode1(headA, headB *ListNode) *ListNode {
	p1 := headA
	p2 := headB
	for p1 != p2 {
		if p1 == nil {
			p1 = headB
		} else {
			p1 = p1.Next
		}
		if p2 == nil {
			p2 = headA
		} else {
			p2 = p2.Next
		}
	}
	return p1
}

/*
官方的方法一也可以想到的，拿空间换时间。
官方通过遍历一个表，将其转为map，再遍历另一个表查，是O(m+n)
我认为可以同时遍历和查询，两个表的内容存储在同一个表不冲突。
*/

func getIntersectionNode2(headA, headB *ListNode) *ListNode {
	m := make(map[*ListNode]bool)
	p1 := headA
	p2 := headB
	for p1 != nil || p2 != nil {
		if _, ok := m[p1]; ok {
			return p1
		} else if p1 != nil {
			m[p1] = true
			p1 = p1.Next
		}
		if _, ok := m[p2]; ok {
			return p2
		} else if p2 != nil {
			m[p2] = true
			p2 = p2.Next
		}
	}
	return nil
}
