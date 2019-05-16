/*
 * @lc app=leetcode id=112 lang=golang
 *
 * [112] Path Sum
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 测试集 [1,2] 1
func hasPathSum(root *TreeNode, sum int) bool {
	//DFS 递归实现
	if root == nil{
		return false
	}
	if root.Val == sum && root.Left == nil && root.Right == nil{ //必须是叶子节点
		return true
	}
	return hasPathSum(root.Left,sum-root.Val) || hasPathSum(root.Right, sum-root.Val)
}

