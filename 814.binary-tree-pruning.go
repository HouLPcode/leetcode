/*
 * @lc app=leetcode id=814 lang=golang
 *
 * [814] Binary Tree Pruning
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
//  (0 ms)
func pruneTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	if root.Left != nil {
		pruneTree(root.Left)
	}
	if root.Right != nil {
		pruneTree(root.Right)
	}

	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil && root.Left.Val == 0 {
		root.Left = nil
	}
	if root.Right != nil && root.Right.Left == nil && root.Right.Right == nil && root.Right.Val == 0 {
		root.Right = nil
	}
	// 典型错误，root = nil并不能影响父节点的指向
	// if root.Left == nil && root.Right == nil && root.Val == 0{
	//     root = nil
	// }

	return root
}
