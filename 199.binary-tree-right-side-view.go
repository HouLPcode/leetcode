/*
 * @lc app=leetcode id=199 lang=golang
 *
 * [199] Binary Tree Right Side View
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rightSideView(root *TreeNode) []int {
	//返回每行最右侧的节点
	// BFS, 非递归
	if root == nil{
		return []int{}
	}
	rnt := []int{}
	queue := list.New()
	queue.PushBack(root)
	for queue.Front() != nil {
		//先把最右侧的节点记录下来
		rnt = append(rnt, queue.Back().Value.(*TreeNode).Val)
		l := queue.Len()
		for i:=0; i<l; i++{
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
	return rnt
}

