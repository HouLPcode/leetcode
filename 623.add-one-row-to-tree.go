/*
 * @lc app=leetcode id=623 lang=golang
 *
 * [623] Add One Row to Tree
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// (44 ms) √ Your runtime beats 94.12 %
func addOneRow(root *TreeNode, v int, d int) *TreeNode {
	// if root == nil { 题目要求非空
	//     return root
	// }
	if d == 1 {
		return &TreeNode{
			Val:   v,
			Left:  root,
			Right: nil,
		}
	}
	// 找到第d-1层的节点
	queue := []*TreeNode{}
	queue = append(queue, root)
	for len(queue) > 0 {
		d--
		if d == 1 {
			break
		}
		size := len(queue)
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}
	for i := 0; i < len(queue); i++ {
		queue[i].Left = &TreeNode{v, queue[i].Left, nil}
		queue[i].Right = &TreeNode{v, nil, queue[i].Right}
	}
	return root
}
