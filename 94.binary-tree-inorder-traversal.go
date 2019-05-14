/*
 * @lc app=leetcode id=94 lang=golang
 *
 * [94] Binary Tree Inorder Traversal
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 // 递归实现
// func inorderTraversal(root *TreeNode) []int {
// 	// 怎么把节点值存储在[]int 中
// 	if root == nil{
// 		return []int{}
// 	}
// 	rnt := []int{} 
// 	lefts := inorderTraversal(root.Left)
// 	rights := inorderTraversal(root.Right)
// 	rnt = append(rnt, lefts...)
// 	rnt = append(rnt, root.Val)
// 	rnt = append(rnt, rights...)
// 	return rnt
// }

// 非递归，
func inorderTraversal(root *TreeNode) []int {
	// 怎么把节点值存储在[]int 中
	if root == nil{
		return []int{}
	}
	stack := []*TreeNode{} // 初始空栈，不需要root入栈
	rnt := []int{} 
	for len(stack) > 0 || root != nil{ // 注意此处不能只判断栈空
		if root != nil{ // 注意此处不能是循环
			stack = append(stack,root)
			root = root.Left
		}else{
			// 出栈
			root,stack = stack[len(stack)-1],stack[:len(stack)-1]
			// 访问根节点
			rnt = append(rnt,root.Val)
			// 指向右节点
			root = root.Right
		}
	}
	return rnt
}

