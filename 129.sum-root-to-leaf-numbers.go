/*
 * @lc app=leetcode id=129 lang=golang
 *
 * [129] Sum Root to Leaf Numbers
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumNumbers(root *TreeNode) int {
	// DFS
	total := 0
	dfsNum(root,0,&total)
	return total
}

func dfsNum(root *TreeNode, num int, total *int){
	if root == nil{
		return
	}
	num = num*10 + root.Val
	if root.Left == nil && root.Right == nil{
		*total += num
		return 
	}
	dfsNum(root.Left, num, total)
	dfsNum(root.Right, num, total)
}
