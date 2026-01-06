package LeetCodeHot100

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
	if head.Next == nil { //链表中节点数目在范围[1, 105] 内
		return true
	} else if head.Next.Next == nil {
		if head.Val == head.Next.Val {
			return true
		} else {
			return false
		}
	}
	slow, fast := head.Next, head.Next.Next //单数 ：slow 中位数 fast 回文尾
	newSlow := slow.Next                    //双数 ：slow、newSlow 中位数 fast.Next
	slow.Next = head
	head.Next = nil
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		p := slow
		slow = newSlow
		newSlow = slow.Next
		slow.Next = p
	}
	var p1 *ListNode
	p2 := newSlow
	if fast.Next == nil {
		p1 = slow.Next
	} else if fast.Next.Next == nil {
		p1 = slow
	}
	for p1 != nil {
		if p1.Val != p2.Val {
			return false
		}
		p1 = p1.Next
		p2 = p2.Next
	}
	return true
}

/*
【题解】
很容易想到转存数组，从中位数开始检查，但这需要O(n)的空间
快慢指针是很好的O(n)时间O(1)空间的解法，而且我认为我比官方题解好的地方是：我反转的是前半个数组，从中间向两头比较；官方反转的是后半个数组，从两头向中间比较。
官方的做法仍旧需要先遍历完一遍数组，再比较，O(n+n/2)，好处是反转指针不会有断裂处；
n1>n2>n3>n4>n5  =>  n1>n2>n3<n4<n5
我的做法可以在找中位数的过程中就完成反转指针，比较只是走完后半程，O(n)，但反转指针会出现断裂处，需要小心处理。
n1>n2>n3>n4>n5  =>  n1<n2<n3 n4>n5
*/
/*
剩下的数组和递归的方法就先不记录了，不是很难
*/
