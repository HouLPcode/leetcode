/*
 * @lc app=leetcode id=307 lang=golang
 *
 * [307] Range Sum Query - Mutable
 */
// 52 ms, faster than 61.63%
// segment tree 实现， 用来计算[s,e]的区间和
// 主要操作： build， update， rangeQuery
// ["NumArray"]\n[[[]]]
// ["NumArray","update","sumRange","sumRange","update","sumRange"]\n[[[9,-8]],[0,3],[1,1],[0,1],[1,-3],[0,1]]

type NumArray struct {
	Node *SegNode
}

func Constructor(nums []int) NumArray {
	return NumArray{
		build(0, len(nums)-1, nums),
	}
}

func (this *NumArray) Update(i int, val int) {
	this.Node.update(i, val)
}

func (this *NumArray) SumRange(i int, j int) int { // query的过程
	return this.Node.queryRange(i, j)
}

type SegNode struct {
	Left, Right *SegNode // 左右节点
	Start, End  int      // 区间
	Sum         int      // 和，也可以是区间中的最值
}

// build O(n)
func build(s, e int, vals []int) *SegNode {
	if len(vals) == 0 {
		return &SegNode{}
	}
	if s == e {
		return &SegNode{
			Left:  nil,
			Right: nil,
			Sum:   vals[s],
			Start: s,
			End:   e,
		}
	}
	mid := (e-s)/2 + s
	l := build(s, mid, vals)
	r := build(mid+1, e, vals)
	return &SegNode{
		Left:  l,
		Right: r,
		Sum:   l.Sum + r.Sum,
		Start: s,
		End:   e,
	}
}

// O(log(n))
func (root *SegNode) update(i, val int) {
	if root.Start == i && root.End == i {
		root.Sum = val
		return
	}
	mid := (root.End-root.Start)/2 + root.Start
	if i <= mid {
		root.Left.update(i, val)
	} else if i > mid {
		root.Right.update(i, val)
	}
	// 注意，经过上面递归，孩子节点的Sum已经更新，此处需要更新该节点的Sum
	root.Sum = root.Left.Sum + root.Right.Sum
}

func (root *SegNode) queryRange(s, e int) int {
	if root.Start == s && root.End == e {
		return root.Sum
	}
	mid := (root.End-root.Start)/2 + root.Start
	if s > mid {
		return root.Right.queryRange(s, e)
	} else if e <= mid { // 注意，左子树包含mid节点
		return root.Left.queryRange(s, e)
	}
	return root.Left.queryRange(s, mid) + root.Right.queryRange(mid+1, e)
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * obj.Update(i,val);
 * param_2 := obj.SumRange(i,j);
 */
