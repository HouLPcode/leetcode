/*
 * @lc app=leetcode id=572 lang=golang
 *
 * [572] Subtree of Another Tree
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 20 ms, faster than 62.26%
// [3,4,5,1,2,null,null,0]\n[4,1,2]
// [1,1]\n[1] 注意，可以有重复元素
// [3,4,5,1,null,2]\n[3,1,2]
// [1,null,1,null,1,null,1,null,1,null,1,null,1,null,1,null,1,null,1,null,1,2]\n[1,null,1,null,1,null,1,null,1,null,1,2]
func isSubtree(s *TreeNode, t *TreeNode) bool {
	if s == nil && t == nil {
		return true
	}
	if s == nil || t == nil {
		return false
	}
	b := isSame(s, t)
	if b == true {
		return true
	}
	return isSubtree(s.Left, t) || isSubtree(s.Right, t)
}

func isSame(s *TreeNode, t *TreeNode) bool {
	if s == nil && t == nil {
		return true
	}
	if s == nil || t == nil {
		return false
	}
	return (s.Val == t.Val) && isSame(s.Left, t.Left) && isSame(s.Right, t.Right)
}
