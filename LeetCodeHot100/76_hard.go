package LeetCodeHot100

func minWindow(s string, t string) string {
	var res string
	ori, cur := make(map[string]int), make(map[string]int)
	oriUp, curUp, oriLow, curLow := 1<<26, 1<<26, 1<<26, 1<<26
	for i := range t {
		oriUp, oriLow = addStringToInt(t[i:i+1], oriUp, oriLow)
		ori[t[i:i+1]]++
	}
	var index []int //下标队列
	var l, r, p int
	for i := range s {
		if ori[s[i:i+1]] > 0 {
			cur[s[i:i+1]]++
			if cur[s[i:i+1]] >= ori[s[i:i+1]] {
				curUp, curLow = addStringToInt(s[i:i+1], curUp, curLow)
			}
			l = i
			r = i + 1
			break
		}
	}
	if oriUp+oriLow == curUp+curLow { //首位字符就是字串
		return t
	}
	for r < len(s) {
		v := s[r : r+1]
		if n, ok := ori[v]; ok {
			index = append(index, r)
			cur[v]++
			if cur[v] >= n {
				curUp, curLow = addStringToInt(v, curUp, curLow)
				for oriUp+oriLow == curUp+curLow {
					res = minString(res, s[l:r+1])
					v = s[l : l+1]
					l = index[p]
					p++
					cur[v]--
					if cur[v] < ori[v] {
						curUp, curLow = deleteStringToInt(v, curUp, curLow)
					}
				}
			}
		}
		r++
	}
	return res
}

func minString(a, b string) string { // 别忘了初始是空值
	if len(a) == 0 {
		return b
	} else if len(b) == 0 {
		return a
	}
	if len(a) > len(b) {
		return b
	} else {
		return a
	}
}

func addStringToInt(s string, up, low int) (int, int) {
	if s[0] < 'a' {
		up |= 1 << (s[0] - 'A')
	} else {
		low |= 1 << (s[0] - 'a')
	}
	return up, low
}

func deleteStringToInt(s string, up, low int) (int, int) {
	if s[0] < 'a' {
		up ^= 1 << (s[0] - 'A')
	} else {
		low ^= 1 << (s[0] - 'a')
	}
	return up, low
}

/*
【题解】（int26位数组太麻烦，大小写很烦，写map吧）
滑动窗口，记录最小字串。维护两个map（一个对照，一个累计），对应元素位置增增减减，同时维护2个序列（一个对照，一个累计）
快一点的话，再存一个下标队列，先进先出
移动，元素是否存在？先确定起始位置
右移，元素是否存在？
	不存在继续；
	存在，
		1）下标入队，
		2）累计map对应+1，
			2.1）当个数>=对照对应时，序列对应位为1，并立即判断两个序列是否相等，
				2.1.1）当序列相等时，判断记录字串；下标出队，左移并删减累计map对应，若无效删除，则重新判断记录字串，若有效删除，序列对应位为0，右移循环这一模块
思路其实是清晰的，但我按思路写出来竟然花了将近两个小时，除了马虎的地方，基本是不需要大修的，但是写的内容太多，作为常规功能是不易理解的。
官解写的虽然清晰但我看不上，我认为效率是不如我的。我的与做加法、异或做减法的思想还是很棒的。对比序列数比遍历对比map好多了。
我多写了一个下标用于快速收缩左窗口，而官解的做法是一点一点收缩对比，有点蠢，他还遍历map对比，好低的效率。我也可以一点点收缩（没写），我的对比效率高。
*/
