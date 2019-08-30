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
 */
// 0 ms
// [1]
// [1,2,null,4,5,6,]
// BFS 层次遍历，直到遇到第一个空节点后，队列中应该全是空节点
func isCompleteTree(root *TreeNode) bool {
	if root == nil {
		return true
	}

	queue := []*TreeNode{}
	queue = append(queue, root)
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node == nil { // 直到找到第一个空节点
			break
		}
		queue = append(queue, node.Left, node.Right)
	}
	// 空节点之后，应该全是空节点
	for _, node := range queue {
		if node != nil {
			return false
		}
	}
	return true
}
