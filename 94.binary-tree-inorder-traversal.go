/*
 * @lc app=leetcode id=94 lang=golang
 *
 * [94] Binary Tree Inorder Traversal
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 // 递归实现
func inorderTraversal(root *TreeNode) []int {
	// 怎么把节点值存储在[]int 中
	if root == nil{
		return []int{}
	}
	rnt := []int{} 
	lefts := inorderTraversal(root.Left)
	rights := inorderTraversal(root.Right)
	rnt = append(rnt, lefts...)
	rnt = append(rnt, root.Val)
	rnt = append(rnt, rights...)
	return rnt
}

