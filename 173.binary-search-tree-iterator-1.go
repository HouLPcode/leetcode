/*
 * @lc app=leetcode id=173 lang=golang
 *
 * [173] Binary Search Tree Iterator
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
//  44 ms, faster than 77.78%
// 缓存所有节点
type BSTIterator struct {
	Buf   []int
	Index int
}

func Constructor(root *TreeNode) BSTIterator {
	// 中序遍历，缓存数据
	rtn := BSTIterator{}
	inOrder(root, &rtn.Buf)
	rtn.Index = 0
	return rtn
}

func inOrder(root *TreeNode, seq *[]int) {
	if root == nil {
		return
	}
	inOrder(root.Left, seq)
	*seq = append(*seq, root.Val)
	inOrder(root.Right, seq)
}

/** @return the next smallest number */
func (this *BSTIterator) Next() int {
	this.Index++
	return this.Buf[this.Index-1]
}

/** @return whether we have a next smallest number */
func (this *BSTIterator) HasNext() bool {
	return this.Index < len(this.Buf)
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */