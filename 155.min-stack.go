/*
 * @lc app=leetcode id=155 lang=golang
 *
 * [155] Min Stack
 */
 // 创建一个辅助栈空间，维持栈的最小元素
 // 入栈值比当前值小，则该值入栈，否则最小值入栈
type MinStack struct {
	Data []int
	Min []int
}


/** initialize your data structure here. */
func Constructor() MinStack {
    return MinStack{
		[]int{},
		[]int{},
	}
}

func (this *MinStack) Push(x int)  {
	this.Data = append(this.Data,x)
	if len(this.Min) == 0{
		this.Min = append(this.Min,x)
	}else if this.Min[len(this.Min)-1] > x{ // 非空
			this.Min = append(this.Min,x)
	}else{
			this.Min = append(this.Min,this.Min[len(this.Min)-1])
	}
}

// 不考虑空pop和空top
func (this *MinStack) Pop()  {
    this.Data = this.Data[:len(this.Data)-1]
    this.Min = this.Min[:len(this.Min)-1]
}


func (this *MinStack) Top() int {
    return this.Data[len(this.Data)-1]
}


func (this *MinStack) GetMin() int {
    return this.Min[len(this.Min)-1]
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */

