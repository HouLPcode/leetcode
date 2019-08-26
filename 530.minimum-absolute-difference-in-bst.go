import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 12 ms, faster than 80.95%
// [236,104,701,null,227,null,911]
// 注意是任意两节点，不一定是父子节点之间，有可能是根节点和右子树最左节点，左子树最右节点
func getMinimumDifference(root *TreeNode) int {
	if root == nil {
		return math.MaxInt64
	}
	// 计算root节点的差值
	sub := math.MaxInt64
	if root.Left != nil {
		p := root.Left
		for p.Right != nil {
			p = p.Right
		}
		sub = root.Val - p.Val
	}
	if root.Right != nil {
		p := root.Right
		for p.Left != nil {
			p = p.Left
		}
		if p.Val-root.Val < sub {
			sub = p.Val - root.Val
		}
	}
	return min(sub, getMinimumDifference(root.Left), getMinimumDifference(root.Right))
}

func min(a, b, c int) int {
	if a <= b && a <= c {
		return a
	} else if b <= a && b <= c {
		return b
	}
	return c
}
