/*
 * @lc app=leetcode id=1026 lang=golang
 *
 * [1026] Maximum Difference Between Node and Ancestor
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 相当于找tree的最大最小节点，
// 一定注意，题目中要求两个最值节点不能在同一层。最大最小值不能在同一层
func maxAncestorDiff(root *TreeNode) int {
	// 题目中给出节点值得范围 0-100000.
	minnum := 100000
	maxnum := 0
	return getmaxAncestorDiff(root, minnum, maxnum)
}

// root为根节点，当前已知的最大最小值，返回最大差值
// 怎么避免的最大最小值不在同一层？？？ DFS，访问的是路径上的最值，一条路径上不可能同时获取同一层的最值
func getmaxAncestorDiff(root *TreeNode, minnum,maxnum int) int{
	if root == nil {
		return maxnum - minnum
	}
	// 访问当前节点
	minnum = min(minnum, root.Val)
	maxnum = max(maxnum, root.Val)
	// 递归左右子树
	leftMaxDiff := getmaxAncestorDiff(root.Left, minnum, maxnum) // 左子树最大差值
	rightMaxDiff := getmaxAncestorDiff(root.Right, minnum, maxnum) // 右子树最大差值
	// 返回最大数值差
	return max(leftMaxDiff, rightMaxDiff)
} 

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

