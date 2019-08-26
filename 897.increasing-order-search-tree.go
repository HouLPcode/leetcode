/*
 * @lc app=leetcode id=897 lang=golang
 *
 * [897] Increasing Order Search Tree
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// (24 ms) √ Your runtime beats 54.09 %
func increasingBST(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	if root.Left == nil {
		root.Left = nil
		root.Right = increasingBST(root.Right)
		return root
	}

	n := increasingBST(root.Left)
	p := n
	for p.Right != nil {
		p = p.Right
	}
	p.Right = root
	root.Left = nil
	root.Right = increasingBST(root.Right)
	return n
}
