import "math"

/*
 * @lc app=leetcode id=783 lang=golang
 *
 * [783] Minimum Distance Between BST Nodes
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
//  0 ms
// minimum difference是最小差值
// 注意不一定是直接父子节点
// [90,69,null,49,89,null,52,null,null,null,null]
// 中序遍历，当前节点减去前节点值，找到最小的
// 注意，特出处理下最左节点，他没有前节点，如果设置成最小值，最小差值初始为最大值，两者作差比较会出错
func minDiffInBST(root *TreeNode) int {
	rtn := math.MaxInt64
	pre := math.MinInt64
	inOrder(root, &pre, &rtn)
	return rtn
}

// pre表示前一个节点
func inOrder(root *TreeNode, pre, min *int) {
	if root == nil {
		return
	}
	inOrder(root.Left, pre, min)
	if *pre == math.MinInt64 {
		// 最左节点，没有前节点 ？？？？？？优化
	} else if root.Val-*pre < *min { // 当前节点肯定大于前节点
		*min = root.Val - *pre
	}
	*pre = root.Val
	inOrder(root.Right, pre, min)
}

