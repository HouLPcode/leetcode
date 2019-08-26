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
// 12 ms, faster than 75.71%
func findTilt(root *TreeNode) int {
	if root == nil {
		return 0
	}
	t, _ := help(root)
	return t
}

//
func help(root *TreeNode) (tile, sum int) {
	if root == nil {
		return 0, 0
	}
	tl, sl := help(root.Left)
	tr, sr := help(root.Right)
	if sl > sr {
		return tl + tr + sl - sr, root.Val + sl + sr
	}
	return tl + tr + sr - sl, root.Val + sl + sr
}
