/*
 * @lc app=leetcode id=102 lang=golang
 *
 * [102] Binary Tree Level Order Traversal
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
	//怎么区分一层,
	// 一次处理一层,每层生成一个[]slice，之后其添加到[][]
	if root == nil{
		return [][]int{}
	}
	out := [][]int{}
	list := list.New()
	list.PushBack(root)
	for list.Front()!=nil{
		levelLen := list.Len()
		currentLevel := []int{}
		for i :=0;i<levelLen;i++{
			node := list.Front()
			list.Remove(node)
			currentLevel = append(currentLevel,node.Value.(*TreeNode).Val)
			if node.Value.(*TreeNode).Left != nil{
				list.PushBack(node.Value.(*TreeNode).Left)
			}
			if node.Value.(*TreeNode).Right != nil{
				list.PushBack(node.Value.(*TreeNode).Right)
			}
		}
		out = append(out,currentLevel)
	}
	return out
}

