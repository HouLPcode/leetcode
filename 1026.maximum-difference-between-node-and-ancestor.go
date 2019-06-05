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
// 相当于找tree的最大最小节点，注意最大最小值不能在同一层
func maxAncestorDiff(root *TreeNode) int {
	if root == nil {
		return 0
	}
	max, min := root.Val, root.Val
	queue := list.New()
	queue.PushBack(root)
	for queue.Front() != nil {
		node := queue.Front()
		queue.Remove(node)
		if node.Value.(*TreeNode).Val > max {
			max = node.Value.(*TreeNode).Val
		}
		if node.Value.(*TreeNode).Val < min {
			min = node.Value.(*TreeNode).Val
		}
		if node.Value.(*TreeNode).Left != nil{
			queue.PushBack(node.Value.(*TreeNode).Left)
		}
		if node.Value.(*TreeNode).Right != nil{
			queue.PushBack(node.Value.(*TreeNode).Right)
		}
	}
	return max - min
}
