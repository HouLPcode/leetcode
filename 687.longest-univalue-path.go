/*
 * @lc app=leetcode id=687 lang=golang
 *
 * [687] Longest Univalue Path
 */
// 72 ms, faster than 86.11%
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func longestUnivaluePath(root *TreeNode) int {
	ans := 0
	univaluePath(root, &ans)
	return ans
}

// 计算当前节点相同值的最大单边数
// 计算的过程中统计当前已知的最大边数
func univaluePath(root *TreeNode, maxnum *int) int {
	if root == nil {
		return 0
	}
	// 统计左右子树的最大同值边数
	lPaths := univaluePath(root.Left, maxnum) // 此处没有直接初始成0在if中计算是为了主动调用，更新max值
	rPaths := univaluePath(root.Right, maxnum)
	if root.Left != nil && root.Left.Val == root.Val {
		lPaths++ // 算上当前节点的边
	} else {
		lPaths = 0 // 一定注意此处清零，表示包括该节点的单边是0
	}
	if root.Right != nil && root.Right.Val == root.Val {
		rPaths++
	} else {
		rPaths = 0
	}
	if *maxnum < lPaths+rPaths {
		*maxnum = lPaths + rPaths // 更新已知最大边数
	}

	if lPaths > rPaths {
		return lPaths
	}
	return rPaths
}