/*
 * @lc app=leetcode id=226 lang=golang
 *
 * [226] Invert Binary Tree
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func invertTree(root *TreeNode) *TreeNode {
	// 递归实现
	if root == nil{
		return nil
	}
	// 典型错误，必须同时赋值
	// root.Left = invertTree(root.Right)
	// root.Right = invertTree(root.Left)
	// return root

	root.Left, root.Right = invertTree(root.Right), invertTree(root.Left)
	return root
}

