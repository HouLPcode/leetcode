/*
 * @lc app=leetcode id=530 lang=golang
 *
 * [530] Minimum Absolute Difference in BST
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
//  [236,104,701,null,227,null,911]
// 典型思路错误：所有节点的最小差值不是根和左右子节点的差值
func getMinimumDifference(root *TreeNode) int {
    if root == nil {
		return math.MaxInt64 //注意，不合法的调用返回的是int最大值
	}
	sub := math.MaxInt64
	if root.Left != nil {
		sub = root.Val - root.Left.Val
	}
	if root.Right != nil {
		if root.Right.Val - root.Val < sub {
			sub = root.Right.Val - root.Val
		}
	}
	return min(sub, getMinimumDifference(root.Left), getMinimumDifference(root.Right))
}

func min(a, b, c int) int {
	if a <=b && a <= c {
		return a 
	}else if b <= a && b <= c {
		return b
	}
	return c
}

