/*
 * @lc app=leetcode id=988 lang=golang
 *
 * [988] Smallest String Starting From Leaf
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 典型错误
// [4,0,1,1]
// 递归
func smallestFromLeaf(root *TreeNode) string {
	if root == nil {
		return "|" // 这是一个比 'z'大的字符，表示放弃这个选项
	}

	if root.Left == nil && root.Right == nil {
		return string(byte(root.Val) + 'a') // 叶子节点返回的是val字符
	}

	lstr := smallestFromLeaf(root.Left)
	rstr := smallestFromLeaf(root.Right)
	// 典型错误 1，l和r中最小的加上root值后不一定还是最小的
	return min(lstr, rstr) + string(byte(root.Val)+'a')

	// 典型错误 2  注意，此处一定要先拼接root，然后再比较
	return min(lstr+string(byte(root.Val)+'a'), rstr+string(byte(root.Val)+'a'))
}

func min(a, b string) string {
	if a < b {
		return a
	}
	return b
}
