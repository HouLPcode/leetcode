/*
 * @lc app=leetcode id=865 lang=golang
 *
 * [865] Smallest Subtree with all the Deepest Nodes
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// (0 ms)
func subtreeWithAllDeepest(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	ldep := depth(root.Left)
	rdep := depth(root.Right)
	if ldep == rdep {
		return root
	} else if ldep > rdep {
		return subtreeWithAllDeepest(root.Left)
	}
	return subtreeWithAllDeepest(root.Right)
}

func depth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l := depth(root.Left)
	r := depth(root.Right)
	return max(l, r) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
