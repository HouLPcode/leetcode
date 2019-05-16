/*
 * @lc app=leetcode id=513 lang=golang
 *
 * [513] Find Bottom Left Tree Value
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findBottomLeftValue(root *TreeNode) int {
	// BFS 扫描一遍
	//题目要求root != nil
	var rtn int
	queue := list.New()
	queue.PushBack(root)
	for queue.Front() != nil{
		l := queue.Len()
		rtn = queue.Front().Value.(*TreeNode).Val
		for i:=0;i<l;i++{
			node := queue.Front()
			queue.Remove(node)
			if node.Value.(*TreeNode).Left != nil{
				queue.PushBack(node.Value.(*TreeNode).Left)
			}
			if node.Value.(*TreeNode).Right != nil{
				queue.PushBack(node.Value.(*TreeNode).Right)
			}
		}
	}
	return rtn
}
