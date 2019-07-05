/*
 * @lc app=leetcode id=637 lang=golang
 *
 * [637] Average of Levels in Binary Tree
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
//  12 ms  98.99 %
func averageOfLevels(root *TreeNode) []float64 {
	// 题目给出root不为空
	rtn := []float64{}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		sum := 0
		l := len(queue)
		for i := l; i > 0; i-- {
			node := queue[0]
			queue = queue[1:]
			sum += node.Val
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		rtn = append(rtn, float64(sum)/float64(l))
	}
	return rtn
}

