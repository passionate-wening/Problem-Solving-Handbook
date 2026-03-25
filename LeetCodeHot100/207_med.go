package LeetCodeHot100

func canFinish(numCourses int, prerequisites [][]int) bool {
	courses := make([][]int, numCourses)
	for _, v := range prerequisites {
		courses[v[0]] = append(courses[v[0]], v[1])
	}
	visited := make([]int, numCourses)
	flag := true
	var find func(key int)
	find = func(key int) {
		visited[key] = 1
		for _, v := range courses[key] {
			if visited[v] == 1 {
				flag = false
				return
			} else if visited[v] == 0 {
				find(v)
				if !flag {
					return
				}
			}
		}
		visited[key] = 2
	}
	for i := 0; i < numCourses; i++ {
		if visited[i] == 0 {
			find(i)
		}
		if !flag {
			return flag
		}
	}
	return flag
}

/*
【题解】
历史想法：维护1个map，快速查找对应关系，纪录被访问过的节点（可以用负数记录），只要确保单个叶子向上回溯无环即可。（存在多个祖先）多用些内存吧，写分离一些：记录所有节点，每个节点单独检查是否有环。
考虑树的形态，先验课程是当前课程的多杈分枝，只要这个分支不构成环就可以。但我们只写过二叉树。
考虑倒过来，某先验课程到后续课程的路径是不应该有环的，所以可以倒着找一下。如果没环，就裁枝，看下一个先验。
所以需要先map映射先验和后续。
注意：1）当前课程不能是自己的先验课程，这是回环；2）这不是树，这是有向图，是可以闭口的，比如4是1、2的先验，而1、2都是3的先验
不行不行不行，思路不好，不要再试错了

官解：拓扑排序————就是要用有向图来解，要学会有向图的代码逻辑
1）有向无环图
2）拓扑排序不止一种
请先完成210题的学习————拓扑排序
经典拓扑排序：三位状态+dfs
果然，学完210，再做这道题都不用动脑子，就写对了...
*/
