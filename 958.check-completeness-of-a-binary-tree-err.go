/*
 * @lc app=leetcode id=958 lang=golang
 *
 * [958] Check Completeness of a Binary Tree
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */ err-------------------------------------
// [1]
// [1,2,null,4,5,6]
// BFS 层次遍历，统计node个数，直到遇到第一个空节点为止
func isCompleteTree(root *TreeNode) bool {
	if root == nil {
		return true
	}

	queue := []*TreeNode{}
	queue = append(queue, root)
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node.Left == nil || node.Right == nil {
			if node.Left == nil && node.Right != nil {
				return false
			}
			break
		}
		queue = append(queue, node.Left, node.Right)
	}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node.Left != nil || node.Right != nil {
			return false
		}
	}

	return true
}
