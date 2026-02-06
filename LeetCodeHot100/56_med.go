package LeetCodeHot100

func merge(intervals [][]int) [][]int {
	for i := 0; i < len(intervals); i++ {
		for j := i + 1; j < len(intervals); j++ {
			if intervals[j][0] < intervals[i][0] {
				temp := intervals[i]
				intervals[i] = intervals[j]
				intervals[j] = temp
			}
		}
	}
	var res [][]int
	for _, in := range intervals {
		if len(res) == 0 || res[len(res)-1][1] < in[0] {
			res = append(res, in)
		} else {
			res[len(res)-1][1] = max(res[len(res)-1][1], in[1])
		}
	}
	return res
}

/*
【题解】
维护一组map，遍历
- 区间在范围外，跳过(头>val || 尾<key)
- 区间有重合(头>key&&尾>val)(头<key&&尾<val>)
	- 修改map的key（取小）/val（取大）
- 区间包含(头<key&&尾>val)(头>key&&尾<val)
	- 删掉被包含的，如果没有重合，就新增

好的，做错了，错的放下面。看了题解，先排序：这样就能保证前半部分不会重合；后合并/新增
*/

func mergeXXXXXXXX(intervals [][]int) [][]int {
	m := make(map[int]int)
	for _, i := range intervals {
		flag := true
		for key, val := range m {
			if i[0] > val || i[1] < key {
			} else if i[0] >= key && i[1] >= val {
				m[key] = i[1]
				i[0] = key
				flag = false
			} else if i[0] <= key && i[1] <= val {
				delete(m, key)
				m[i[0]] = val
				i[1] = val
				flag = false
			} else if i[0] <= key && i[1] >= val {
				delete(m, key)
			} else if i[0] >= key && i[1] <= val {
				flag = false
			}
		}
		if flag {
			m[i[0]] = i[1]
		}
	}
	var res [][]int
	for k, v := range m {
		res = append(res, []int{k, v})
	}
	return res
}
