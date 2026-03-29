package LeetCodeHot100

import (
	"fmt"
	"strconv"
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

func TestMoveZeroes(t *testing.T) {
	n := []int{0, 1, 0, 3, 12}
	moveZeroes1(n)
	fmt.Println(n)
	n = []int{0}
	moveZeroes1(n)
	fmt.Println(n)
}

func TestIsPalindrome(t *testing.T) {
	n := []int{1, 3, 5, 3, 1}
	head := &ListNode{Val: 0}
	p := head
	for _, v := range n {
		p.Next = &ListNode{Val: v}
		p = p.Next
	}
	isPalindrome(head.Next)
}

func TestDetectCycle(t *testing.T) {
	n := []int{3, 2, 0, -4}
	head := &ListNode{Val: 0}
	p := head
	for _, v := range n {
		p.Next = &ListNode{Val: v}
		p = p.Next
	}
	p.Next = head.Next.Next
	detectCycle1(head.Next)
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

func TestLongestConsecutive(t *testing.T) {
	fmt.Println(longestConsecutive1([]int{1, 0, 1, 2}))
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

func TestThreeSum(t *testing.T) {
	threeSum([]int{2, -3, 0, -2, -5, -5, -4, 1, 2, -2, 2, 0, 2, -4, 5, 5, -10})
}

func TestLengthOfLongestSubstring(t *testing.T) {
	fmt.Println(lengthOfLongestSubstring1("pwwkew"))
	fmt.Println(lengthOfLongestSubstring1("bbbbb"))
	fmt.Println(lengthOfLongestSubstring1("aab"))
	fmt.Println(lengthOfLongestSubstring1(""))
}

func TestFindAnagrams(t *testing.T) {
	//fmt.Println(findAnagrams2("abab", "ab"))
	fmt.Println(findAnagrams2("bpaa", "aa"))
}

func TestSubarraySum(t *testing.T) {
	fmt.Println(subarraySum([]int{1, 2, -3}, 0))
}

func TestFlatten(t *testing.T) {
	var temp []*TreeNode
	temp = append(temp, &TreeNode{Val: 1})
	temp = append(temp, &TreeNode{Val: 2})
	fmt.Println(len(temp))
	temp = temp[:len(temp)-1]
	fmt.Println(len(temp))
}

func TestMaxSlidingWindow(t *testing.T) {
	a := maxSlidingWindow([]int{1, -1}, 1)
	fmt.Println(a)
}

func TestSetZeroes(t *testing.T) {
	m := [][]int{{1, 0, 1}, {1, 0, 1}, {1, 1, 1}}
	setZeroes1(m)
	fmt.Println(m)
}

func TestSortList(t *testing.T) {
	m := []int{4, 2, 1, 3}
	sortList2(buildListTest(m))
}

func buildListTest(nums []int) *ListNode {
	list := &ListNode{Val: 0}
	p := list
	for _, v := range nums {
		node := &ListNode{Val: v}
		p.Next = node
		p = p.Next
	}
	return list.Next
}

func TestSpiralOrder(t *testing.T) {
	spiralOrder([][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}})
	spiralOrder([][]int{})
	spiralOrder([][]int{{1}})
}

func TestSearchMatrix(t *testing.T) {
	searchMatrix([][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 13)
}

func TestMerge(t *testing.T) {
	x := merge([][]int{{5, 5}, {1, 3}, {3, 5}, {4, 6}, {1, 1}, {3, 3}, {5, 6}, {3, 3}, {2, 4}, {0, 0}})
	fmt.Println(x)
}
func TestMinStack(t *testing.T) {
	s := ConstructorMinStack()
	s.Push(-10)
	s.Push(14)
	s.Push(-20)
	s.Pop()
	s.Push(10)
	s.Push(-7)
	s.Push(-7)
	s.Pop()
	s.Pop()
}

func TestRotate(t *testing.T) {
	rotate2([]int{-1, -100, 3, 99}, 2)
}

func TestRotate48(t *testing.T) {
	rotate_48([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}})
}

func TestPathSum(t *testing.T) {
	root := buildBTreeTest([]string{"5", "4", "8", "11", "null", "13", "4", "7", "2", "null", "null", "5", "1"})
	pathSum(root, 22)
}

func buildBTreeTest(nums []string) *TreeNode {
	if len(nums) == 0 {
		return &TreeNode{}
	}
	n, e := strconv.Atoi(nums[0])
	root := &TreeNode{}
	queen := []*TreeNode{}
	if e == nil {
		root = &TreeNode{Val: n}
		queen = append(queen, root)
	}
	i := 0
	for in := 0; in < len(queen); in++ {
		p := queen[in]
		if i+1 < len(nums) {
			n, e = strconv.Atoi(nums[i+1])
			if e == nil {
				left := &TreeNode{Val: n}
				queen = append(queen, left)
				p.Left = left
			}
			i++
		}
		if i+1 < len(nums) {
			n, e = strconv.Atoi(nums[i+1])
			if e == nil {
				right := &TreeNode{Val: n}
				queen = append(queen, right)
				p.Right = right
			}
			i++
		}
	}
	return root
}

func TestOrangesRotting(t *testing.T) {
	orangesRotting1([][]int{{2, 1, 1}, {1, 1, 1}, {0, 1, 2}})
}

//func TestConstructoranFinish(t *testing.T) {
//	canFinish(2, [][]int{{2, 1}, {1, 0}})
//}

func TestDecodeString(t *testing.T) {
	decodeString("3[a2[c]]")
}

func TestDailyTemperatures(t *testing.T) {
	dailyTemperatures1([]int{73, 74, 75, 71, 69, 72, 76, 73, 73})
}

func TestPermute(t *testing.T) {
	//permute([]int{1, 2, 3})
}

func TestSearch(t *testing.T) {
	res := search1([]int{4, 5, 6, 7, 0, 1, 2}, 0)
	fmt.Println(res)
}

func TestLargestRectangleArea(t *testing.T) {
	//largestRectangleArea([]int{3, 6, 5, 7, 4, 8, 1, 0})
}

func TestMinWindow(t *testing.T) {
	minWindow("abc", "aabc")
}

func TestCombinationSum(t *testing.T) {
	combinationSum([]int{2, 3, 6, 7}, 7)
}

func TestWordBreak(t *testing.T) {
	wordBreak("leetcode", []string{"leet", "code"})
}

func TestCanPartition(t *testing.T) {
	canPartition1([]int{1, 5, 11, 5})
}

func TestNextPermutation(t *testing.T) {
	nextPermutation([]int{3, 2, 1})
}

func TestUniquePaths(t *testing.T) {
	uniquePaths(3, 2)
}

func TestExist(t *testing.T) {
	exist([][]byte{[]byte("ABCE"), []byte("SFES"), []byte("ADEE")}, "ABCESEEEFS")
}

func TestCanFinish(t *testing.T) {
	canFinish(3, [][]int{{1, 0}, {2, 0}, {0, 2}})
}

func TestFindOrder(t *testing.T) {
	findOrder(1, [][]int{})
}

func TestPartition(t *testing.T) {
	partition("cdd")
}

func TestMinPathSum(t *testing.T) {
	minPathSum([][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}})
}

func TestFindKthLargest(t *testing.T) {
	findKthLargest([]int{3, 2, 1, 5, 6, 4}, 2)
}

func TestLongestCommonSubsequence1(t *testing.T) {
	longestCommonSubsequence1("abcde", "ace")
}

func TestSolveNQueens(t *testing.T) {
	solveNQueens(4)
}
