/*
 * @lc app=leetcode id=307 lang=golang
 *
 * [307] Range Sum Query - Mutable
 */
//  (44 ms) √ Your runtime beats 94.19 %
// FenwickTree / binary index tree
type NumArray struct {
	Tree []int
}

func Constructor(nums []int) NumArray {
	rtn := NumArray{
		make([]int, len(nums)+1), // 从1开始
	}
	for k, v := range nums {
		rtn.update(k+1, v) // 从1开始
	}
	return rtn
}

func (this *NumArray) Update(i int, val int) {
	if i == 0 {
		val = val - this.query(i+1)
	} else {
		val = val - (this.query(i+1) - this.query(i))
	}
	this.update(i+1, val)
}

func (this *NumArray) SumRange(i int, j int) int { // query的过程
	return this.query(j+1) - this.query(i) // 包含[i,j]
}

// 自己添加的
func (this *NumArray) update(i int, val int) {
	for i < len(this.Tree) {
		this.Tree[i] += (val)
		i += (i & (-i)) // 加上末尾1
	}
}

func (this NumArray) query(i int) int {
	sum := 0
	for i > 0 {
		sum += this.Tree[i]
		i -= (i & (-i))
	}
	return sum
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * obj.Update(i,val);
 * param_2 := obj.SumRange(i,j);
 */
