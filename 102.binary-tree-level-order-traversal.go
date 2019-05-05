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
	//怎么区分一层,两个队列交替使用实现分层？？？
	out := [][]int{}
	if root == nil{
		return out
	}
	list1 := list.New()
	list2 := list.New()
	list1.PushBack(root)
	i:=0
	out = append(out,[]int{})
	out[i] = append(out[i],root.Left.Val)
	i++
	for {
		out = append(out,[]int{})
		for list1.Len() != 0{
			node := list1.Front()
			list1.Remove(node)
			if node.Value.(*TreeNode).Left != nil{
				out[i] = append(out[i],node.Value.(*TreeNode).Left.Val)
				list2.PushBack(node.Value.(*TreeNode).Left)
			}
			if node.Value.(*TreeNode).Right != nil{
				out[i] = append(out[i],node.Value.(*TreeNode).Right.Val)
				list2.PushBack(node.Value.(*TreeNode).Right)
			}
		}
		i++
		out = append(out,[]int{})
		for list2.Len() != 0{
			node := list2.Front()
			list2.Remove(node)
			if node.Value.(*TreeNode).Left != nil{
				out[i] = append(out[i],node.Value.(*TreeNode).Left.Val)
				list1.PushBack(node.Value.(*TreeNode).Left)
			}
			if node.Value.(*TreeNode).Right != nil{
				out[i] = append(out[i],node.Value.(*TreeNode).Right.Val)
				list1.PushBack(node.Value.(*TreeNode).Right)
			}
		}	
		i++
	}
	return out
}

