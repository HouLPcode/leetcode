/*
 * @lc app=leetcode id=103 lang=golang
 *
 * [103] Binary Tree Zigzag Level Order Traversal
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 典型错误，不能一次遍历从左往右，下一次从右向左，因为出队的序列决定过了子节点在整层遍历中的序列
func zigzagLevelOrder(root *TreeNode) [][]int {
	rnt := [][]int{}
	
	// 队列
	queue := list.New()
	// 入队
	queue.PushBack(root)
	for queue.Front() != nil{
		tmp := []int{}
		for node := queue.Front(); node != nil;{
			queue.Remove(node)//出队
			tmp = append(tmp, node.(TreeNode).Val)
			if node.(*TreeNode).Right != nil{
				queue.PushBack(node.(*TreeNode).Right)
			}
			if node.(*TreeNode).Left != nil{
				queue.PushBack(node.(*TreeNode).Left)
			}
		}
		rnt = append(rnt,tmp)
		tmp = []int{}
		for node := queue.Front(); node != nil;{
			queue.Remove(node)//出队
			tmp = append(tmp,node.(*TreeNode).Val)
			if node.Left != nil{
				queue.PushBack(node.Left)
			}
			if node.Right != nil{
				queue.PushBack(node.Right)
			}	
		}
		rnt = append(rnt,tmp)
	}
	return rnt
}

