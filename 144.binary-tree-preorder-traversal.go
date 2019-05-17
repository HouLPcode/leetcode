/*
 * @lc app=leetcode id=144 lang=golang
 *
 * [144] Binary Tree Preorder Traversal
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func preorderTraversal(root *TreeNode) []int {
	// 递归
	if root == nil{
		return []int{}
	}
	rtn := []int{}
	lnodes := preorderTraversal(root.Left)
	rnodes := preorderTraversal(root.Right)
	rtn = append(rtn,root.Val)
	rtn = append(rtn, lnodes...)
	rtn = append(rtn, rnodes...)
	return rtn
}

