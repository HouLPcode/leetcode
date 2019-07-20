import "math"

/*
 * @lc app=leetcode id=124 lang=golang
 *
 * [124] Binary Tree Maximum Path Sum
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
//  (20 ms) 76.92 %
//  [-3]
func maxPathSum(root *TreeNode) int {
	rtn := math.MinInt64 // 一定注意，最终结果可能是负值，此处不能初始化为0
	pathSum(root, &rtn)
	return rtn
}

// 返回单边最大路径和
func pathSum(root *TreeNode, maxPath *int) int {
	if root == nil {
		return math.MinInt64 // 此处返回int最小值
	}
	lPaths := pathSum(root.Left, maxPath)
	rPaths := pathSum(root.Right, maxPath)
	sum := root.Val
	if lPaths > 0 {
		sum += lPaths
	} else {
		lPaths = 0 // 如果子树最大路径和是负值，则值返回当前节点值，不计算子树
	}
	if rPaths > 0 {
		sum += rPaths
	} else {
		rPaths = 0
	}

	if *maxPath < sum {
		*maxPath = sum
	}
	// 注意，返回的是单边的最大边数
	if lPaths > rPaths {
		return root.Val + lPaths
	}
	return root.Val + rPaths
}
