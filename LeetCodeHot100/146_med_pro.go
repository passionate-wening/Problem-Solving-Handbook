package LeetCodeHot100

type CacheNode struct {
	Prev, Next *CacheNode
	Key, Value int //key和Value都要存，方便删除时知道是哪个key!!!
}

type LRUCache struct {
	cache      map[int]*CacheNode
	capacity   int
	Len        int
	head, tail *CacheNode
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		cache:    make(map[int]*CacheNode),
		capacity: capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	cache, ok := this.cache[key]
	if !ok {
		return -1
	} else {
		this.Update(cache)
		return cache.Value
	}
}

func (this *LRUCache) Put(key int, value int) {
	cache, ok := this.cache[key]
	if ok {
		cache.Value = value
		this.Update(cache)
	} else {
		cache = &CacheNode{Key: key, Value: value}
		this.PushBack(cache)
		if this.Len > this.capacity {
			delete(this.cache, this.head.Key) //不要忘记删cache
			this.Remove(this.head)
		}
	}
	this.cache[key] = cache //及时更新指针
}

// 对于已有元素，更新位置，cache不处理，指针跟随变化
func (this *LRUCache) Update(node *CacheNode) {
	if this.tail == node {
		return
	} else {
		this.Remove(node)
		this.PushBack(node)
	}
}

// 根据规则head为空时，tail一定为空；
// 注意：我们拿到的可能是历史节点，所以要把Next清空
func (this *LRUCache) PushBack(node *CacheNode) {
	if this.head == nil {
		this.head = node
		this.tail = node
	} else {
		this.tail.Next = node
		node.Prev = this.tail
		node.Next = nil
		this.tail = this.tail.Next
	}
	this.Len++
}

// 由于capacity不为0，head和tail就不可能指向同一个
func (this *LRUCache) Remove(node *CacheNode) {
	if this.head == node {
		this.head = this.head.Next
	} else if this.tail == node {
		this.tail = this.tail.Prev
	} else {
		node.Prev.Next = node.Next
		node.Next.Prev = node.Prev
	}
	node.Next = nil
	node.Prev = nil
	this.Len--
}
