package LeetCodeHot100

func findOrder(numCourses int, prerequisites [][]int) []int {
	m := make([][]int, numCourses)
	for _, v := range prerequisites {
		m[v[0]] = append(m[v[0]], v[1])
	}
	st := []int{}
	status := make([]int, numCourses) //0 1 2
	flag := false
	var find func(key int)
	find = func(key int) {
		status[key] = 1
		list := m[key]
		for _, v := range list {
			if status[v] == 0 {
				find(v)
				if flag {
					return
				}
			} else if status[v] == 1 {
				flag = true
				return
			}
		}
		status[key] = 2
		st = append(st, key) //mark
	}
	for i := 0; i < numCourses; i++ { //mark
		if status[i] == 0 {
			find(i)
		}
	}
	if flag {
		return []int{}
	}
	return st
}

/*
【题解】为了学习，直接看的题解。分别记录深度优先和广度优先两种。
[深度优先]：关注两个点 1）不能有环 2）拓扑排序有多种，任选一种即可
所以题解给出了一个节点的三种状态：未搜索/搜索中/已完成
由于学习课程是先验关系，所以排序符合栈的思想，后进先出，这样全部搜索完成，栈顶就是不需要先验课程的课程。
对于无环的判断是单次搜索要进行的，所以状态只能将“未搜索”状态转为“搜索中”，回溯时改成“已完成”，不然就是有环。
| 「未搜索」：从来没有搜索过这个节点；
| 「搜索中」：正在搜索这个节点，由单次搜索更新，若更新前该节点已经为此状态，说明有环；
| 「已完成」：历史已经搜索过并且回溯过这个节点了，即该节点已经入栈，并且所有该节点的相邻节点都出现在栈的更底部的位置，满足拓扑排序的要求。
要遍历课程啊！
*/
