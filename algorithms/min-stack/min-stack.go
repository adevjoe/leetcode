// https://leetcode.com/problems/min-stack

package leetcode

type MinStack struct {
	head *Stack
}

type Stack struct {
	val  int
	min  int
	next *Stack
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	head := &Stack{
		val:  val,
		next: this.head,
	}
	if this.head == nil || val < this.GetMin() {
		head.min = val
	} else {
		head.min = this.GetMin()
	}
	this.head = head
}

func (this *MinStack) Pop() {
	this.head = this.head.next
}

func (this *MinStack) Top() int {
	return this.head.val
}

func (this *MinStack) GetMin() int {
	return this.head.min
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
