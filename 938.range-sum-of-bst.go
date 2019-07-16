/*
 * @lc app=leetcode id=938 lang=golang
 *
 * [938] Range Sum of BST
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
//  92 ms, faster than 92.28%
// 递归
func rangeSumBST(root *TreeNode, L int, R int) int {
	// 中序遍历，递归 / 非递归
	if root == nil {
		return 0
	}
	sum := 0
	inOrder(root, L, R, &sum)
	return sum
}

func inOrder(root *TreeNode, L, R int, sum *int) {
	if root == nil {
		return
	}

	if root.Val >= L { //剪枝，不满足条件的节点不再递归调用
		inOrder(root.Left, L, R, sum)
	}

	if root.Val >= L && root.Val <= R { // 此处必须有判断，上面一层判断只是过滤掉小于L的子节点，但是节点的值可能会大于R，需要舍去。同理，下面的判断舍去了大于R的节点，但是会存在小于L的子节点
		(*sum) += root.Val
	}

	if root.Val <= R {
		inOrder(root.Right, L, R, sum)
	}
}
