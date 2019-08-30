/*
 * @lc app=leetcode id=998 lang=golang
 *
 * [998] Maximum Binary Tree II
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 0 ms
func insertIntoMaxTree(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{val, nil, nil}
	}
	if val > root.Val {
		rtn := &TreeNode{}
		rtn.Val = val
		rtn.Left = root
		return rtn
	}
	// 新的节点插入右子树，同时注意树的连接，不能断开
	root.Right = insertIntoMaxTree(root.Right, val)
	return root
}
