/*
 * @lc app=leetcode id=104 lang=golang
 *
 * [104] Maximum Depth of Binary Tree
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 递归实现
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	ldep := maxDepth(root.Left)
	rdep := maxDepth(root.Right)
	if ldep < rdep {
		return 1 + rdep
	}
	return 1 + ldep
}
