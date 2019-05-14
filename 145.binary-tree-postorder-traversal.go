/*
 * @lc app=leetcode id=145 lang=golang
 *
 * [145] Binary Tree Postorder Traversal
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// func postorderTraversal(root *TreeNode) []int {
// 	// 非递归，后续遍历
// 	if root == nil{
// 		return []int{}
// 	}
// 	rnt := []int{}
// 	lefts := postorderTraversal(root.Left)
// 	rights := postorderTraversal(root.Right)
// 	rnt = append(rnt, lefts...)// 注意...不能少
// 	rnt = append(rnt, rights...)
// 	rnt = append(rnt, root.Val)
// 	return rnt
// }

// 非递归，后序，1个栈
func postorderTraversal(root *TreeNode) []int {
	if root == nil{
		return []int{}
	}
	var top *TreeNode = nil
	stack := []*TreeNode{}
	rnt := []int{}
	// 头节点入栈
	stack = append(stack, root)
	for len(stack) > 0{
		top = stack[len(stack)-1]

		if top.Left!=nil && top.Left!=root && top.Right!=root{
			stack = append(stack, top.Left)
		}else if top.Right!=nil && top.Right!=root{
			stack = append(stack, top.Right)
		}else{
			//访问栈顶元素
			root = top
			// 出栈
			top,stack = stack[len(stack)-1],stack[:len(stack)-1]
			// 访问站定元素
			rnt = append(rnt, top.Val)
		}
	}
	return rnt
}

