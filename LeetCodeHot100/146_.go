package LeetCodeHot100

import (
	"container/list"
)

type LRUCache_ struct {
	data     map[any]int
	queue    *list.List
	queueMap map[any]*list.Element
	capacity int
}

func Constructor_(capacity int) LRUCache_ {
	return LRUCache_{
		data:     make(map[any]int),
		queue:    list.New(),
		queueMap: make(map[any]*list.Element),
		capacity: capacity,
	}
}

func (this *LRUCache_) Get_(key int) int {
	data, ok := this.data[key]
	if !ok {
		return -1
	} else {
		oldKey := this.queueMap[key]
		this.queue.Remove(oldKey)
		newKey := this.queue.PushBack(key)
		this.queueMap[key] = newKey
		return data
	}
}

func (this *LRUCache_) Put_(key int, value int) {
	defer func() {
		if this.queue.Len() > this.capacity {
			oldKey := this.queue.Front()
			delete(this.queueMap, oldKey.Value)
			delete(this.data, oldKey.Value)
			this.queue.Remove(oldKey)
		}
	}()
	_, ok := this.data[key]
	if ok {
		oldKey := this.queueMap[key]
		this.queue.Remove(oldKey)
	}
	this.data[key] = value
	newKey := this.queue.PushBack(key)
	this.queueMap[key] = newKey
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

/*
【题解】注意：要求get、put都为O(1)
我能想到的就是：一个map用于快速查找，一个key的队列/链用于更新缓存（队列设置容量先进先出；链用双指针控制滑动窗口，感觉队列也是链写的）
虽然解出来了，符合官方思路，但是debug了很久，且不符合面试要求
1、存储的内容太多，queueMap其实就是需要的cache，既然有了cache，就没必要单独存键值map了；
2、借用list库是取巧，面试要求自己可以实现双链表！
内容较多，我将单独写文件，146_pro.go为自实现情况，146.go为官方解答（推荐），请多写几遍，至手熟。
*/
