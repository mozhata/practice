package basic

/*
用栈实现队列
https://leetcode-cn.com/problems/implement-queue-using-stacks/
*/

type MyQueue struct {
	stack   []int
	reverse []int
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	// reverse 中的元素倒到stack中
	for len(this.reverse) > 0 {
		val := this.reverse[len(this.reverse)-1]
		this.reverse = this.reverse[:len(this.reverse)-1]
		this.stack = append(this.stack, val)
	}
	this.stack = append(this.stack, x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	// 从reverse里面找,需要先把stack中的内容倒到reverse
	for len(this.stack) > 0 {
		val := this.stack[len(this.stack)-1]
		this.stack = this.stack[:len(this.stack)-1]
		this.reverse = append(this.reverse, val)
	}
	if len(this.reverse) == 0 {
		return -1
	}
	val := this.reverse[len(this.reverse)-1]
	this.reverse = this.reverse[:len(this.reverse)-1]
	return val
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	// 从reverse里面找,需要先把stack中的内容倒到reverse
	for len(this.stack) > 0 {
		val := this.stack[len(this.stack)-1]
		this.stack = this.stack[:len(this.stack)-1]
		this.reverse = append(this.reverse, val)
	}
	if len(this.reverse) == 0 {
		return -1
	}
	return this.reverse[len(this.reverse)-1]
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return len(this.reverse) == 0 && len(this.stack) == 0
}
