/*
 * @lc app=leetcode id=1145 lang=golang
 *
 * [1145] Binary Tree Coloring Game
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 0 ms
// 二手着色的最佳位置是 x的父节点，左，右子节点中节点数最多的那个
func btreeGameWinningMove(root *TreeNode, n int, x int) bool {
	// find x
	nodeX := findnode(root, x)
	lNodes := nodes(nodeX.Left)
	rNodes := nodes(nodeX.Right)
	pNodes := n - lNodes - rNodes - 1
	return max3(lNodes, rNodes, pNodes) > n/2
}

// TODO find 和 计算节点个数，两个函数可以合并
func findnode(root *TreeNode, x int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == x {
		return root
	}
	if node := findnode(root.Left, x); node != nil {
		return node
	}
	return findnode(root.Right, x)
}

// 返回root为根的树的节点个数
func nodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + nodes(root.Left) + nodes(root.Right)
}

func max3(a, b, c int) int {
	if a >= b && a >= c {
		return a
	} else if b >= a && b >= c {
		return b
	}
	return c
}
