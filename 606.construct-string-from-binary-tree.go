import "strconv"

/*
 * @lc app=leetcode id=606 lang=golang
 *
 * [606] Construct String from Binary Tree
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 12ms 59.65%
// 前序递归调用
func tree2str(t *TreeNode) string {
	if t == nil {
		return "" // nil 返回空字符串
	}
	rtn := strconv.Itoa(t.Val)
	if t.Left == nil && t.Right == nil {
		return strconv.Itoa(t.Val) // 孩子节点，直接返回
	}
	if t.Right == nil {
		return strconv.Itoa(t.Val) + "(" + tree2str(t.Left) + ")"
	}
	return strconv.Itoa(t.Val) + "(" + tree2str(t.Left) + ")" +
		"(" + tree2str(t.Left) + ")" // 不管左节点是不是空，都加一个（）
}
