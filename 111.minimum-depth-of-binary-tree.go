/*
 * @lc app=leetcode id=111 lang=golang
 *
 * [111] Minimum Depth of Binary Tree
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
func minDepth(root *TreeNode) int {
    if root == nil{
		return 0
	}
	if root.Left == nil || root.Right == nil{
		return 1 + minDepth(root.Right) + minDepth(root.Left)
	}

	ldep := minDepth(root.Left)
	rdep := minDepth(root.Right)
	if ldep < rdep {
		return 1 + ldep
	}
	return 1 + rdep
}

