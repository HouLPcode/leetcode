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
    rnt := []int{} 
	inorder(root,&rnt)
	return rnt
}

func inorder(root *TreeNode, rnt *[]int){
	if root == nil{
		return 
	}
	inorder(root.Left, rnt)
	*rnt = append(*rnt,root.Val)
	inorder(root.Right, rnt)
}

