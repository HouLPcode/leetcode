/*
 * @lc app=leetcode id=979 lang=golang
 *
 * [979] Distribute Coins in Binary Tree
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
//  (0 ms)
func distributeCoins(root *TreeNode) int {
	rtn := 0
	moves(root, &rtn)
	return rtn
}

// root树平衡需要从外面获取的节点数，可能正，可能负
func moves(root *TreeNode, ans *int) int {
	if root == nil {
		return 0
	}

	l := moves(root.Left, ans)
	r := moves(root.Right, ans)
	*ans += abs(l) + abs(r) // 移动的数目和root.Val 没有关系
	return root.Val - 1 + l + r
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
