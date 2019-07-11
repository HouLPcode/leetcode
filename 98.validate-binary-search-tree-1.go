import "math"

/*
 * @lc app=leetcode id=98 lang=golang
 *
 * [98] Validate Binary Search Tree
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 4 ms 98.17 %
func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return root.Val > getMax(root.Left) &&
		root.Val < getMin(root.Right) &&
		isValidBST(root.Left) &&
		isValidBST(root.Right)
}

func getMin(root *TreeNode) int {
	if root == nil {
		return math.MaxInt64 // 注意此处返回的值正好相反，保证调用该函数处为true
	}
	for root.Left != nil {
		root = root.Left
	}
	return root.Val
}

func getMax(root *TreeNode) int {
	if root == nil {
		return math.MinInt64
	}
	for root.Right != nil {
		root = root.Right
	}
	return root.Val
}
