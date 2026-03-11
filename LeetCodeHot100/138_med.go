package LeetCodeHot100

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
	m := make(map[*Node]*Node)
	root := &Node{Val: 0}
	p := root
	for head != nil {
		n, ok := m[head]
		if !ok {
			n = &Node{Val: head.Val}
			m[head] = n
		}
		p.Next = n
		p = p.Next
		random, ok := m[head.Random]
		if !ok && head.Random != nil {
			random = &Node{Val: head.Random.Val}
			m[head.Random] = random
		}
		p.Random = random
		head = head.Next
	}
	return root.Next
}

/*
【题解】
简单的，维护一个新旧指针表就行了。可以先存一遍表，也可以边存边取，像我这样。
*/
