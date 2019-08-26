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
// [3,4,5,1,2,null,null,0]\n[4,1,2]
// [1,1]\n[1] 注意，可以有重复元素
// [3,4,5,1,null,2]\n[3,1,2]
func isSubtree(s *TreeNode, t *TreeNode) bool {
	if s == nil && t == nil {
		return true
	}
	if s == nil || t == nil {
		return false
	}

	// 典型错误 一定注意，此处是或的关系
	return ((s.Val == t.Val) && isSubtree(s.Left, t.Left) && isSubtree(s.Right, t.Right)) ||
		isSubtree(s.Left, t) ||
		isSubtree(s.Right, t)

	// 典型错误
	// 	if s.Val == t.Val {
	// 		return isSubtree(s.Left, t.Left) && isSubtree(s.Right, t.Right)
	// }
	// return isSubtree(s.Left, t) || isSubtree(s.Right, t)
}
