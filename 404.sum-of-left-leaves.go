/*
 * @lc app=leetcode id=404 lang=golang
 *
 * [404] Sum of Left Leaves
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
//  递归遍历每个节点，同时找该节点的直接左叶子节点
func sumOfLeftLeaves(root *TreeNode) int {
    if root == nil {
		return 0
	}
	if root.Left != nil && 
		root.Left.Left==nil && 
		root.Left.Right == nil {
			return root.Left.Val + sumOfLeftLeaves(root.Left) + sumOfLeftLeaves(root.Right)
		}
	return sumOfLeftLeaves(root.Left) + sumOfLeftLeaves(root.Right)
}

