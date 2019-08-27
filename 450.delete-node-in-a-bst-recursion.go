/*
 * @lc app=leetcode id=450 lang=golang
 *
 * [450] Delete Node in a BST
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
//  316 ms, faster than 34.57%
// 递归方式
// 只处理删除root节点的情况
// root有左右子树，将右子树最小值替换为root，然后再次递归删除最小值节点
func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	// if root.Val < key {
	// 	return deleteNode(root.Right, key) // 一定注意，此处不是直接返回，而是更新左右节点
	// } else if root.Val > key {
	// 	return deleteNode(root.Left, key)
	// }
	if root.Val < key {
		root.Right = deleteNode(root.Right, key)
	} else if root.Val > key {
		root.Left = deleteNode(root.Left, key)
	} else {
		// 删除根节点
		if root.Left == nil {
			return root.Right
		} else if root.Right == nil {
			return root.Left
		}
		// 找右子树最左节点
		p := root.Right
		for p.Left != nil {
			p = p.Left
		}
		root.Val = p.Val
		root.Right = deleteNode(root.Right, p.Val)
	}

	return root
}
