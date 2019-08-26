/*
 * @lc app=leetcode id=563 lang=golang
 *
 * [563] Binary Tree Tilt
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 16 ms, faster than 31.43%
func findTilt(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l := sum(root.Left)
	r := sum(root.Right)
	if l > r {
		return l - r + findTilt(root.Left) + findTilt(root.Right)
	}
	return r - l + findTilt(root.Left) + findTilt(root.Right)
}

func sum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return root.Val + sum(root.Left) + sum(root.Right)
}
