/*
 * @lc app=leetcode id=303 lang=golang
 *
 * [303] Range Sum Query - Immutable
 */

 //题目要求，nums数组不能改变
 //NumArray初始化的时候可以将NumArray[i]表示成前i项的和
 // 时间O(1) 空间O(n)
type NumArray struct {
	nums []int
	datas []int
}


func Constructor(nums []int) NumArray {
	datas := make([]int,len(nums),len(nums))
	for k,v := range nums{
		if k == 0{ // 怎么设计能不破坏for流水线？？？
			datas[k] = v
			continue
		}
		datas[k] = datas[k-1] + v 
	}
	return NumArray{nums,datas}
}


func (this *NumArray) SumRange(i int, j int) int {
	return this.datas[j] - this.datas[i] + this.nums[i]
}


/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(i,j);
 */

