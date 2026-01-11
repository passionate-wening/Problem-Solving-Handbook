package LeetCodeHot100

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverse(head, tail *ListNode) {
	pre := &ListNode{Val: 0, Next: head}
	p1, p2 := pre, pre.Next
	for p1 != tail { //即p2是尾之后的节点
		p3 := p2.Next
		p2.Next = p1
		p1 = p2
		p2 = p3
	}
	pre.Next.Next = nil
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	preHead := &ListNode{Val: 0, Next: head}
	temp := preHead
	for temp.Next != nil {
		pHead, pTail := temp.Next, temp
		i := 0
		for ; i < k && pTail.Next != nil; i++ {
			pTail = pTail.Next
		}
		if i < k {
			break
		}
		node := pTail.Next    //暂存新尾后继，逆转后这个指针就断了
		reverse(pHead, pTail) //（此时头节点被清空了后继指针）
		temp.Next = pTail     //尾是新头，前驱节点指向新头
		pHead.Next = node     //头是新尾，头节点指向新尾后继
		temp = pHead          //新尾就是下一个分组的前驱节点
	}
	return preHead.Next
}

/*
【题解】
第一遍做理解错题意了，要求的k的意思是分组，不是只搞前k个。
把逆序k个的函数抽出来 ———— 只关注这部分链，头节点最后指向空，参考206题。
需要单独遍历一遍，因为不够分组的部分是不能动的 ———— 确定每一组的头节点和尾节点，默认函数执行完：头是尾，尾是头。要把前后衔接串好。
可以的！困难题也很好的做出来了！
*/
