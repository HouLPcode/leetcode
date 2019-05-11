/*
 * @lc app=leetcode id=100 lang=golang
 *
 * [100] Same Tree
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 // 没有采用尾递归
func isSameTree(p *TreeNode, q *TreeNode) bool {
    if p == nil && q == nil{// 结束条件1
		return true
	}else if p!=nil && q!=nil && p.Val == q.Val{ // nil.Val会报错
		return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
	}
	return false//结束条件2
}

