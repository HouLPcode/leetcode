/*
 * @lc app=leetcode id=701 lang=golang
 *
 * [701] Insert into a Binary Search Tree
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
//  (520 ms) √ Your runtime beats 58.18 %, 不用优化
// 非递归实现
func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{
			Val:   val,
			Left:  nil,
			Right: nil,
		}
	}
	for cur := root; cur != nil; {
		if val < cur.Val {
			if cur.Left == nil {
				cur.Left = &TreeNode{
					Val:   val,
					Left:  nil,
					Right: nil,
				}
				break
			} else {
				cur = cur.Left
			}
		} else {
			if cur.Right == nil {
				cur.Right = &TreeNode{
					Val:   val,
					Left:  nil,
					Right: nil,
				}
				break
			} else {
				cur = cur.Right
			}
		}
	}

	return root
}
