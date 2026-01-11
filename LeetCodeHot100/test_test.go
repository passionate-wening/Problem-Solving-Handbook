package LeetCodeHot100

import (
	"fmt"
	"testing"
)

func TestTwoSum(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9
	res := twoSum1(nums, target)
	fmt.Println(res)
}

func TestMergeTwoLists(t *testing.T) {
	n1 := []int{1, 2, 4}
	n2 := []int{1, 3, 4}
	list1 := &ListNode{Val: 0}
	list2 := &ListNode{Val: 0}
	p1 := list1
	p2 := list2
	for _, v := range n1 {
		node := &ListNode{Val: v}
		p1.Next = node
		p1 = p1.Next
	}
	for _, v := range n2 {
		node := &ListNode{Val: v}
		p2.Next = node
		p2 = p2.Next
	}
	mergeTwoLists(list1.Next, list2.Next)
}

func TestGetIntersectionNode(t *testing.T) {
	n2 := []int{4, 1}
	n1 := []int{5, 6, 1}
	list1 := &ListNode{Val: 0}
	list2 := &ListNode{Val: 0}
	p1 := list1
	p2 := list2
	for _, v := range n1 {
		node := &ListNode{Val: v}
		p1.Next = node
		p1 = p1.Next
	}
	for _, v := range n2 {
		node := &ListNode{Val: v}
		p2.Next = node
		p2 = p2.Next
	}
	n3 := []int{8, 4, 5}
	for _, v := range n3 {
		node := &ListNode{Val: v}
		p1.Next = node
		p1 = p1.Next
		p2.Next = node
		p2 = p2.Next
	}
	getIntersectionNode1(list1.Next, list2.Next)
}

func TestLRUCache(t *testing.T) {
	lRUCache := Constructor(2)
	lRUCache.Put(1, 0) // 缓存是 {1=1}
	printX(lRUCache.head)
	lRUCache.Put(2, 2) // 缓存是 {1=1, 2=2}
	printX(lRUCache.head)
	fmt.Print(lRUCache.Get(1), "\t") // 返回 1
	printX(lRUCache.head)
	lRUCache.Put(3, 3) // 该操作会使得关键字 2 作废，缓存是 {1=1, 3=3}
	printX(lRUCache.head)
	fmt.Print(lRUCache.Get(2), "\t") // 返回 -1 (未找到)
	printX(lRUCache.head)
	lRUCache.Put(4, 4) // 该操作会使得关键字 1 作废，缓存是 {3=3, 4=4}
	printX(lRUCache.head)
	fmt.Print(lRUCache.Get(1), "\t") // 返回 -1 (未找到)
	printX(lRUCache.head)
	fmt.Print(lRUCache.Get(3), "\t") // 返回 3
	printX(lRUCache.head)
	fmt.Print(lRUCache.Get(4), "\t") // 返回 4
	printX(lRUCache.head)
}

func TestSwapPairs(t *testing.T) {
	list := swapPairs1(arrayToList([]int{1, 2, 3, 4}))
	printList(list)
}

func TestReverseKGroup(t *testing.T) {
	list := reverseKGroup(arrayToList([]int{1, 2, 3, 4, 5}), 4)
	printList(list)
}

func TestGroupAnagrams(t *testing.T) {
	fmt.Println(groupAnagrams1([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
}

func arrayToList(nums []int) *ListNode {
	head := &ListNode{Val: 0}
	p1 := head
	for _, v := range nums {
		node := &ListNode{Val: v}
		p1.Next = node
		p1 = p1.Next
	}
	return head.Next
}

func printList(node *ListNode) {
	p := node
	fmt.Print("[")
	for p != nil {
		fmt.Print(p.Val, "\t")
		p = p.Next
	}
	fmt.Println("]")
}

func printX(node *CacheNode) {
	p := node
	fmt.Print("[")
	for p != nil {
		fmt.Print(p.Value, "\t")
		p = p.Next
	}
	fmt.Println("]")
}
