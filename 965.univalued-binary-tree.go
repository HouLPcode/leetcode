/*
 * @lc app=leetcode id=965 lang=golang
 *
 * [965] Univalued Binary Tree
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
func isUnivalTree(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if root.Left != nil && root.Left.Val != root.Val {
		return false
	}
	if root.Right != nil && root.Val != root.Right.Val {
		return false
	}
	return isUnivalTree(root.Left) && isUnivalTree(root.Right)
}
