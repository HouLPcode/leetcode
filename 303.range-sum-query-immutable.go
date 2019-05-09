/*
 * @lc app=leetcode id=303 lang=golang
 *
 * [303] Range Sum Query - Immutable
 */

 //题目要求，nums数组不能改变
 //NumArray初始化的时候可以将NumArray[i]表示成前i项的和
 // 时间O(1) 空间O(n)
type NumArray struct {
	datas []int
}


func Constructor(nums []int) NumArray {
	// i+1存储前i项和，不包含i+1
	datas := make([]int,len(nums)+1,len(nums)+1)
	for k,v := range nums{
		datas[k+1] = datas[k] + v 
	}
	return NumArray{datas}
}


func (this *NumArray) SumRange(i int, j int) int {
	return this.datas[j+1] - this.datas[i]
}


/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(i,j);
 */

