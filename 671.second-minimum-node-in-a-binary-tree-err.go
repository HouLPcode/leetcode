/*
 * @lc app=leetcode id=671 lang=golang
 *
 * [671] Second Minimum Node In a Binary Tree
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
//  [1,1,3,1,1,3,4,3,1,1,1,3,8,4,8,3,3,1,6,2,1]
// BFS, min>root.val err
func findSecondMinimumValue(root *TreeNode) int {
	minnum := root.Val
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node.Val > minnum {
			return node.Val
		}
		if node.Left != nil {
			queue = append(queue, node.Left)
			queue = append(queue, node.Right)
		}
	}
	return -1
}

