/*
 * @lc app=leetcode id=1123 lang=golang
 *
 * [1123] Lowest Common Ancestor of Deepest Leaves
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// (8 ms) √ Your runtime beats 90.91 %
// 返回子树深度相同的节点
func lcaDeepestLeaves(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	ldep := depth(root.Left)
	rdep := depth(root.Right)
	if ldep == rdep {
		return root
	} else if ldep > rdep {
		return lcaDeepestLeaves(root.Left)
	}
	return lcaDeepestLeaves(root.Right)
}

func depth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l := depth(root.Left)
	r := depth(root.Right)
	if l > r {
		return l + 1
	}
	return r + 1
}
