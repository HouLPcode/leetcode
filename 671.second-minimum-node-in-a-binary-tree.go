import "math"

/*
 * @lc app=leetcode id=671 lang=golang
 *
 * [671] Second Minimum Node In a Binary Tree
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
//  [1,1,3,1,1,3,4,3,1,1,1,3,8,4,8,3,3,1,6,2,1]
// DFS, min>root.val err
func findSecondMinimumValue(root *TreeNode) int {
	// non empty
	rtn := math.MaxInt64
	dfs(root, root.Val, &rtn)
	if rtn == math.MaxInt64 {
		return -1
	}
	return rtn
}

func dfs(node *TreeNode, min1 int, min2 *int) {
	if node == nil {
		return
	}
	if node.Val > min1 && node.Val < *min2 {
		*min2 = node.Val
		return
	}
	if node.Val > min1 {
		return // jianzhi
	}
	dfs(node.Left, min1, min2)
	dfs(node.Right, min1, min2)
}

