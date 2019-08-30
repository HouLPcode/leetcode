/*
 * @lc app=leetcode id=337 lang=golang
 *
 * [337] House Robber III
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
//  待优化，吧下面的map去掉，找到dp之间的关系，就能进一步节省空间
func rob(root *TreeNode) int {
	if root == nil {
		return 0
	}
	buf := make(map[*TreeNode]int)
	return robmap(root, buf)
}

func robmap(root *TreeNode, buf map[*TreeNode]int) int {
	if root == nil {
		return 0
	}
	if buf[root] > 0 {
		return buf[root]
	}
	// 打劫root
	val := root.Val
	if root.Left != nil {
		val += robmap(root.Left.Left, buf) + robmap(root.Left.Right, buf)
	}
	if root.Right != nil {
		val += robmap(root.Right.Left, buf) + robmap(root.Right.Right, buf)
	}

	rtn := max(val, robmap(root.Left, buf)+robmap(root.Right, buf))
	buf[root] = rtn
	return rtn // 返回打劫root和不打劫root的最大值
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
