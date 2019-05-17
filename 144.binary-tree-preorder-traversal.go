/*
 * @lc app=leetcode id=144 lang=golang
 *
 * [144] Binary Tree Preorder Traversal
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func preorderTraversal(root *TreeNode) []int {
	// 非递归 ， 栈， 右左入栈
	if root == nil{
		return []int{}
	}
	rtn := []int{}
	stack := []*TreeNode{}//注意栈存的是节点，不是节点的内容
	stack = append(stack, root)
	for len(stack) != 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		rtn = append(rtn, node.Val)
		
		if node.Right != nil{
			stack = append(stack, node.Right)
		}
		if node.Left != nil{
			stack = append(stack, node.Left)
		}
	}
	return rtn
}

