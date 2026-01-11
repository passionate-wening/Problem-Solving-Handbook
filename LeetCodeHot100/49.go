package LeetCodeHot100

func sortAnagrams(strs string) string { //写冒泡吧
	newstrs := []byte(strs)
	for i := 0; i < len(newstrs); i++ {
		for j := i + 1; j < len(newstrs); j++ {
			if newstrs[i] > newstrs[j] {
				temp := newstrs[i]
				newstrs[i] = newstrs[j]
				newstrs[j] = temp
			}
		}
	}
	return string(newstrs)
}

func groupAnagrams(strs []string) [][]string {
	var results [][]string
	for _, s := range strs {
		anagrams := sortAnagrams(s)
		flag := false
		for i, res := range results {
			if res[0] == anagrams {
				results[i] = append(results[i], s)
				flag = true
				break
			}
		}
		if !flag {
			results = append(results, []string{anagrams, s})
		}
	}
	for i, res := range results {
		results[i] = res[1:]
	}
	return results
}

/*
注：字母异位词是通过重新排列不同单词或短语的字母而形成的单词或短语，并使用所有原字母一次。
【题解】
在结果数组中每行第一个写正确顺序的字母排列；	//也可以写个哈希表，顺序数据为键，结果数组行数为值，好处是查询快，坏处是占内存
每个词做字母顺序排序，与结果数组第一个数据比较，有则append，无则新增行；
最后删除每行第一个数据。
！可恶，string不能当数组处理，排不了序！得先转成[]byte数组
注意24和34行：res不是引用值，所以要改results数据具体位置
	101ms	9MB

官方第一个题解写的是存map了，为什么比我写的少，因为它用了官方库的sort，而我自己写的，自己写更优秀！！
（可以背一下官方库的sort写法： sort.Slice(s, func(i,j int) bool { return s[i]<s[j] })）
经过力扣测试用例认证，多占用一点内存，可以提速90%以上，所以map是更好的设计！！！
	6ms	10.40MB
ps：更好的写法，不排序，直接转换26位签名，见最下方
*/

func groupAnagrams1(strs []string) [][]string {
	m := make(map[string][]string)
	for _, s := range strs {
		anagrams := countSign(s) //sortAnagrams(s)	//countSign(s)
		m[anagrams] = append(m[anagrams], s)
	}
	var results [][]string
	for _, group := range m {
		results = append(results, group)
	}
	return results
}

/*
官方题解二也很有意思，map不存排序异位词，存int数组，这个内存消耗更大一点，但不用排序，速度应该更快，试试效果。
	12ms	11.1MB
有趣，并没有更快，也就是这个优化已经没意义了。说实话，异位词数据量大的话，数组挺占内存的。
*/

func groupAnagrams2(strs []string) [][]string {
	m := make(map[[26]int][]string)
	for _, s := range strs {
		var temp [26]int
		for i := 0; i < len(s); i++ {
			temp[s[i]-'a']++
		}
		m[temp] = append(m[temp], s)
	}
	var results [][]string
	for _, group := range m {
		results = append(results, group)
	}
	return results
}

/*
看了下前面网友写的，原来可以不写排序算法，直接依赖特性转换签名 []byte
	9ms	9.96MB
表现不明显，我认为是系统问题了，其实宏观上来看，效果差不多，但这个思路特别好：
不需要占用太多的内存（[]byte数组都是临时的），而且理论上肯定是比排序快的。
*/

func countSign(strs string) string {
	var sign [26]byte
	for i := 0; i < len(strs); i++ {
		sign[strs[i]-'a']++
	}
	return string(sign[:])
}
