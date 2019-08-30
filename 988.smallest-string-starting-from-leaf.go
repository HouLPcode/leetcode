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
//  4 ms, faster than 91.67%
// [4,0,1,1]
// [25,1,null,0,0,1,null,null,null,0]
// 递归
func smallestFromLeaf(root *TreeNode) string {
	return smallTopDown(root, "")
}

func smallTopDown(root *TreeNode, str string) string {
	if root == nil {
		return "|" // 这是一个比 'z'大的字符，表示放弃这个选项
	}

	str = string(byte(root.Val)+'a') + str
	if root.Left == nil && root.Right == nil {
		return str // 叶子节点返回的是val字符
	}

	lstr := smallTopDown(root.Left, str)
	rstr := smallTopDown(root.Right, str)
	// 注意，此处一定要先拼接root，然后再比较
	return min(lstr, rstr)
}

func min(a, b string) string {
	if a < b {
		return a
	}
	return b
}
