package LeetCodeHot100

func maxSlidingWindow1(nums []int, k int) []int {
	var root *ListNode
	for i := 0; i < k; i++ {
		root = updateList(root, nums[i], 0, true)
	}
	res := []int{getHeadVal(root)}
	for i := k; i < len(nums); i++ {
		root = updateList(root, nums[i], nums[i-k], false)
		res = append(res, getHeadVal(root))
	}
	return res
}

func updateList(head *ListNode, insert, delete int, onlyIn bool) *ListNode {
	if head == nil {
		return &ListNode{Val: insert}
	}
	if insert == delete {
		return head
	}
	p := head
	pre := &ListNode{Val: 0, Next: p}
	headPre := pre
	var isIn, isDe bool
	isDe = onlyIn
	for p != nil && (!isIn || !isDe) {
		if !isDe && p.Val == delete {
			pre.Next = p.Next
			p.Next = nil
			p = pre.Next
			isDe = true
			continue
		}
		if !isIn && p.Val < insert {
			in := &ListNode{Val: insert, Next: p}
			pre.Next = in
			isIn = true
		}
		p = p.Next
		pre = pre.Next
	}
	if !isIn { //尾插
		pre.Next = &ListNode{Val: insert}
	}
	return headPre.Next
}

func getHeadVal(head *ListNode) int {
	return head.Val
}

/*
【题解】
可以维护一个窗口大小的逆序链表，这样最坏查k长度。说我超时。
那我们就维护最大值和它的数量？最大值没了再遍历查最大值？
*/

func maxSlidingWindow(nums []int, k int) []int {
	ma, count := findMax(nums[:k])
	res := []int{ma}
	for i := k; i < len(nums); i++ {
		if nums[i-k] == ma {
			count--
		}
		if nums[i] == ma {
			count++
		} else if nums[i] > ma {
			ma = nums[i]
			count = 1
		}
		if count == 0 {
			ma, count = findMax(nums[i-k+1 : i+1])
		}
		res = append(res, ma)
	}
	return res
}

func findMax(nums []int) (int, int) {
	ma, count := nums[0], 0
	for _, n := range nums {
		if ma == n {
			count++
		} else if ma < n {
			ma = n
			count = 1
		}
	}
	return ma, count
}

/*
做对了，但思路不够好，看看官方给出的。
*/
