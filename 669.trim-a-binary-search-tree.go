/*
 * @lc app=leetcode id=669 lang=golang
 *
 * [669] Trim a Binary Search Tree
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
//  8 ms  100 %
//  注意此题递归求解的思路，与BST加减节点算法的比较
func trimBST(root *TreeNode, L int, R int) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val < L {
		return trimBST(root.Right, L, R)
	}
	if root.Val > R {
		return trimBST(root.Left, L, R)
	}
	root.Left = trimBST(root.Left, L, R)
	root.Right = trimBST(root.Right, L, R)
	return root //返回的时候利用root，而不是重新创建一个指针节点，可以加速
}

