/*
 * @lc app=leetcode id=337 lang=golang
 *
 * [337] House Robber III
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// (680 ms) √ Your runtime beats 27.08 %
// 递归， 非DP实现
func rob(root *TreeNode) int {
	if root == nil {
		return 0
	}

	// 打劫root
	val := root.Val
	if root.Left != nil {
		val += rob(root.Left.Left) + rob(root.Left.Right)
	}
	if root.Right != nil {
		val += rob(root.Right.Left) + rob(root.Right.Right)
	}
	return max(val, rob(root.Left)+rob(root.Right)) // 返回打劫root和不打劫root的最大值
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
