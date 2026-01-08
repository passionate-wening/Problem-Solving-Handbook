package LeetCodeHot100

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func checkCycle(head *ListNode) (isCycle bool, a int, b int) {
	if head == nil || head.Next == nil {
		return
	}
	slow, fast := head, head.Next.Next
	for fast != nil && fast.Next != nil {
		a++
		b++
		b++
		if slow == fast {
			return true, a, b
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return
}

func detectCycle(head *ListNode) *ListNode {
	if isCycle, a, b := checkCycle(head); isCycle {
		p1, p2 := head, head
		x := b - a + 1
		for i := 0; i < x; i++ {
			p2 = p2.Next
		}
		for p1 != p2 {
			p1 = p1.Next
			p2 = p2.Next
		}
		return p1
	} else {
		return nil
	}
}

/*
【题解】
先快慢指针确认环，并记录走的步数a、b，那么b-a+1=x（多的步数）；
重新快慢指针，让快指针先走x步，再相遇的就是环的第一个节点。
此方法需要记录步长a、b；
官方题解不需要记录步长。
官方的思路比我更数学一点，先算清楚几个未知数的关系，可以得出：环外的路径a = 第一次相遇后剩下的环路径c+（n-1）*环；
(a+b)*2 = a+b+(b+c)*n  ==>  a=c+(n-1)*(b+c)  &&  环的长度=b+c  &&  slow走的步数=a+b
所以只要slow继续往下走的时候，头部跟着走，就可以在环的第一个节点相遇。
下面给出我根据官方思路写的答案
*/

func detectCycle1(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			p := head
			for slow != p {
				slow = slow.Next
				p = p.Next
			}
			return p
		}
	}
	return nil
}

/*同样，哈希表方法太简单，不记录*/
