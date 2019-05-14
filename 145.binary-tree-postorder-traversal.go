/*
 * @lc app=leetcode id=145 lang=golang
 *
 * [145] Binary Tree Postorder Traversal
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func postorderTraversal(root *TreeNode) []int {
	// 非递归，后续遍历
	if root == nil{
		return []int{}
	}
	rnt := []int{}
	lefts := postorderTraversal(root.Left)
	rights := postorderTraversal(root.Right)
	rnt = append(rnt, lefts...)// 注意...不能少
	rnt = append(rnt, rights...)
	rnt = append(rnt, root.Val)
	return rnt
}

