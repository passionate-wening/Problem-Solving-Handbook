package LeetCodeHot100

type MinStack struct {
	Stack    []int
	Min      int
	MinCount int
}

func ConstructorMinStack() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	this.Stack = append(this.Stack, val)
	if this.MinCount == 0 || this.Min > val {
		this.Min = val
		this.MinCount = 1 //标记一个
	} else if this.Min == val {
		this.MinCount++
	}
}

func (this *MinStack) Pop() {
	if len(this.Stack) == 0 {
		return
	}
	cur := this.Stack[len(this.Stack)-1]
	this.Stack = this.Stack[:len(this.Stack)-1]
	if this.Min == cur {
		this.MinCount--
		if this.MinCount == 0 {
			this.updateMin()
		}
	}
}

func (this *MinStack) updateMin() {
	if len(this.Stack) == 0 {
		this.MinCount = 0
		return
	}
	this.Min = this.Stack[0]
	this.MinCount = 0 //0就是没有数了
	for _, s := range this.Stack {
		if this.Min == s {
			this.MinCount++
		} else if this.Min > s {
			this.Min = s
			this.MinCount = 1 //标记一个
		}
	}
}

func (this *MinStack) Top() int {
	return this.Stack[len(this.Stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.Min
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */

/*
【题解】
根据经验，写数组的性能要比链表好很多。
模仿滑动窗口那个题，维护一个最小元素，当最小元素被删除，则重新遍历查最小（带一个最小数个数也行）
。。。不要马虎。。。
*/
