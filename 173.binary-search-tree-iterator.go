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
// (44 ms)√ Your runtime beats 77.78 %
// 非递归中序遍历的变形
type BSTIterator struct {
	Stack []*TreeNode
}

func Constructor(root *TreeNode) BSTIterator {
	// 中序遍历，缓存数据
	rtn := BSTIterator{
		Stack: []*TreeNode{},
	}
	// 循环将所有左节点入栈
	for p := root; p != nil; p = p.Left {
		rtn.Stack = append(rtn.Stack, p)
	}
	return rtn
}

/** @return the next smallest number */
func (this *BSTIterator) Next() int {
	// top
	node := this.Stack[len(this.Stack)-1]
	// pop
	this.Stack = this.Stack[:len(this.Stack)-1]
	// 左节点循环入栈
	for p := node.Right; p != nil; p = p.Left {
		this.Stack = append(this.Stack, p)
	}
	return node.Val
}

/** @return whether we have a next smallest number */
func (this *BSTIterator) HasNext() bool {
	return len(this.Stack) > 0
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
 