/*
 * @lc app=leetcode id=110 lang=golang
 *
 * [110] Balanced Binary Tree
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 //递归实现
func isBalanced(root *TreeNode) bool {
    if root == nil{
		return true
	}
	// 终止条件的确定,引入高度差
	sub := high(root.Left) - high(root.Right)
	if (sub != 0 && sub != 1 && sub != -1){
		return false
	}
	return isBalanced(root.Left) && isBalanced(root.Right)
}

func high(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + max(high(root.Left), high(root.Right))
}

func max(a,b int) int{
	if a > b{
		return a
	}
	return b
}

