package LeetCodeHot100

type Trie struct {
	children [26]*Trie
	isEnd    bool
}

func ConstructorTrie() Trie {
	return Trie{}
}

func (this *Trie) Insert(word string) {
	cur := this
	for _, w := range word {
		if cur.children[w-'a'] == nil {
			cur.children[w-'a'] = &Trie{}
		}
		cur = cur.children[w-'a']
	}
	cur.isEnd = true
}

func (this *Trie) SearchPrefix(prefix string) *Trie {
	cur := this
	for _, w := range prefix {
		if cur.children[w-'a'] == nil {
			return nil
		}
		cur = cur.children[w-'a']
	}
	return cur
}

func (this *Trie) Search(word string) bool {
	return this.SearchPrefix(word) != nil && this.SearchPrefix(word).isEnd //mark 记得判空
}

func (this *Trie) StartsWith(prefix string) bool {
	return this.SearchPrefix(prefix) != nil
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */

/*
【题解】
笑，我没理解题意，直接map就出来了→_→
Trie（发音类似 "try"）或者说 前缀树 是一种树形数据结构，用于高效地存储和检索字符串数据集中的键。
这一数据结构有相当多的应用情景，例如自动补全和拼写检查。
Trie，又称前缀树或字典树，是一棵有根树，26叉树。。。
 - 指向子节点的指针数组 children。
 - 布尔字段 isEnd 标记结尾。
超简单，主要是在考察你知不知道这个概念，只要知道立马就能写出来。
*/
