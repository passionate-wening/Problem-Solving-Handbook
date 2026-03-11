package LeetCodeHot100

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	res := lists[0]
	for i := 1; i < len(lists); i++ {
		res = mergeTwoList(res, lists[i])
	}
	return res
}

func mergeTwoList(a, b *ListNode) *ListNode {
	res := &ListNode{Val: 0}
	p := res
	for a != nil && b != nil {
		if a.Val > b.Val {
			p.Next = b
			b = b.Next
		} else {
			p.Next = a
			a = a.Next
		}
		p = p.Next
	}
	if a != nil {
		p.Next = a
	}
	if b != nil {
		p.Next = b
	}
	return res.Next
}

/*
【题解】
一起对比有点复杂，最直观的想法就是两两对比。O(NK)，效率不好，不要作为最终答案。熟练mergeTwoList这个函数。
AI面试建议：
	第一步：抛出基础方案（展示思路）
		“这道题最直观的想法是复用‘合并两个有序链表’的逻辑。我可以依次把第 1 个链表和第 2 个合并，结果再和第 3 个合并……直到合并完所有链表。”
		“不过我分析一下复杂度：假设总共有 N 个节点， k 个链表。这种做法每次合并结果链表都会变长，时间复杂度大概是 O(N⋅k)。如果 k 很大，效率会比较低。”
	第二步：提出优化方案（展示深度）
		“为了优化到 O(Nlogk)，我有两种方案：
			1、优先队列（最小堆）：维护 k 个链表的头节点，每次取最小值。
			2、分治法（Divide & Conquer）：类似归并排序，两两配对合并，第一轮合并成 k/2 个链表，第二轮 k/4 个……直到剩 1 个。
		这两种复杂度都是 O(Nlogk)。”
	第三步：根据语言选择实现（展示工程能力）
		“如果是 C++/Java，我会首选优先队列，因为标准库支持很好。但既然我现在用 Go，标准库的 container/heap 需要写不少接口适配代码（样板代码较多）。”
		“为了代码更简洁清晰，我建议使用分治法，它能复用‘合并两个链表’的逻辑，且不需要额外定义堆结构，空间复杂度也更优（递归栈 O(logk) vs 堆 O(k)）。”
		(通常面试官会同意，因为这展示了你对语言特性的理解)
问题一：为什么用分治法，不用优先队列？
	逻辑上优先队列也是O(Nlogk)，但在Go中实现标准库的heap接口需要比较多的样板代码（定义struct、实现Len/Less/Swap/Push/Pop）。分治法代码更简洁，且递归栈空间O(logk)比堆O(k)略优。当然，如果是在C++中，我会毫不犹豫选优先队列。
问题二：分治法的递归深度是多少？
	每次区间减半，深度是O(logk)
问题三：如果链表总数是奇数怎么办？
	分支逻辑里的mid划分自然处理了奇偶情况，最后总会收敛到单个链表。

(有机会手搓下最小堆理解理解吧)
*/

func mergeKLists1(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	return mergeDrive(lists, 0, len(lists)-1)
}

func mergeDrive(lists []*ListNode, left, right int) *ListNode {
	if left > right {
		return nil
	}
	if left == right {
		return lists[left]
	}
	mid := (left + right) / 2
	leftL := mergeDrive(lists, left, mid)
	rightL := mergeDrive(lists, mid+1, right)
	return mergeTwoList(leftL, rightL)
}
