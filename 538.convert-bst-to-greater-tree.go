/*
 * @lc app=leetcode id=538 lang=golang
 *
 * [538] Convert BST to Greater Tree
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
//  188 ms, faster than 86.19%
//  [2,0,3,-4,1]
// [1,0,4,-2,null,3]
func convertBST(root *TreeNode) *TreeNode {
	return help(root, 0)
}

// 比当前节点大的值是右子树的值,如果该节点在左子树中，再加上父节点的值
func help(root *TreeNode, parent int) *TreeNode {
	if root == nil {
		return root
	}

	// help(root.Right, 0) + root.Val?????????
	help(root.Right, parent) // 一定注意，右节点的时候传入的是当前节点的父节点，因为左子树的所有节点小于根节点

	if root.Right != nil { // 注意，有节点可能为空
		// root.Val += root.Right.Val + parent  一定注意，此时的右节点已经加上了parent，不要重复计算
		// root.Val += root.Right.Val // parent
		// 加上的是右节点的最左节点的值
		p := root.Right
		for p.Left != nil {
			p = p.Left
		}
		root.Val += p.Val
	} else {
		root.Val += parent
	}

	help(root.Left, root.Val)

	return root
}
