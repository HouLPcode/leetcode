/*
 * @lc app=leetcode id=225 lang=golang
 *
 * [225] Implement Stack using Queues
 */
 // 两个queue实现
// 入栈：进入非空队列
// 出栈：非空队列转移到另一个队列，只剩一个元素，弹出
// top：类似出栈，获取到元素后将该元素继续入队

import "container/list"//注意包的导入和使用方式

type MyStack struct {
	queue1 *list.List//new出来的是指针类型
	queue2 *list.List
}

/** Initialize your data structure here. */
func Constructor() MyStack {
    return MyStack{
		queue1: list.New(),
		queue2: list.New(),
	}
}


/** Push element x onto stack. */
func (this *MyStack) Push(x int)  {
    if this.queue1.Len() != 0{
		this.queue1.PushBack(x)
	}else{
		this.queue2.PushBack(x)
	}
}


/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {//调用的时候都合法
	rtn := 0
	if this.queue1.Len() > 0{
		for this.queue1.Len() > 1 { //queue1 -> queue2 剩一个元素
			element := this.queue1.Front() //element.Next可以提速？？？
			this.queue2.PushBack(element.Value) // 注意push的是element.Value而不是element，否则会导致类型转换的错误
			this.queue1.Remove(element)
		}
		element := this.queue1.Front()
		this.queue1.Remove(element)
		// rtn = element.Value.(int) interface conversion: interface {} is *list.Element, not int
		rtn = element.Value.(int)//*****一定注意此处的类型断言方式， 
	}else{
		for this.queue2.Len() > 1 { //queue2 -> queue1 剩一个元素
			element := this.queue2.Front()
			this.queue1.PushBack(element.Value)
			this.queue2.Remove(element)
		}
		element := this.queue2.Front()
		this.queue2.Remove(element)
		rtn = element.Value.(int)
	}
	return rtn
}


/** Get the top element. */
func (this *MyStack) Top() int {
    rtn := 0
	if this.queue1.Len() > 0{
		for this.queue1.Len() > 1 { //queue1 -> queue2 剩一个元素
			element := this.queue1.Front()
			this.queue2.PushBack(element.Value)
			this.queue1.Remove(element)
		}
		element := this.queue1.Front()
		rtn= element.Value.(int)
		this.queue2.PushBack(element.Value)
		this.queue1.Remove(element)
	}else{
		for this.queue2.Len() > 1 { //queue2 -> queue1 剩一个元素
			element := this.queue2.Front()
			this.queue1.PushBack(element.Value)
			this.queue2.Remove(element)
		}
		element := this.queue2.Front()
		rtn = element.Value.(int)
		this.queue1.PushBack(element.Value)
		this.queue2.Remove(element)
	}
	return rtn
}


/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
    return this.queue1.Len()==0 && this.queue2.Len()==0
}

// queue的标准操作
// type queue struct{
// }
// func push(){}
// func pop()int{}
// func peek()int{}
// func size()int{}
// func isEmpth(){}


/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */

