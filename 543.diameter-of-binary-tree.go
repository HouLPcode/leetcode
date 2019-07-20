/*
 * @lc app=leetcode id=543 lang=golang
 *
 * [543] Diameter of Binary Tree
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
//  (4 ms)  97.69 %
func diameterOfBinaryTree(root *TreeNode) int {
	rtn := 0
	diameter(root, &rtn)
	return rtn
}

// 返回单边最大边数
func diameter(root *TreeNode, maxPath *int) int {
	if root == nil {
		return 0
	}
	lPaths, rPaths := 0, 0
	if root.Left != nil {
		lPaths = 1 + diameter(root.Left, maxPath)
	}
	if root.Right != nil {
		rPaths = 1 + diameter(root.Right, maxPath)
	}
	if *maxPath < lPaths+rPaths {
		*maxPath = lPaths + rPaths
	}
	// 注意，返回的是单边的最大边数
	if lPaths > rPaths {
		return lPaths
	}
	return rPaths
}
