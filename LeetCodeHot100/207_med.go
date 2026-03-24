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

经典拓扑排序：三位状态+dfs
果然，学完210，再做这道题都不用动脑子，就写对了...
*/
