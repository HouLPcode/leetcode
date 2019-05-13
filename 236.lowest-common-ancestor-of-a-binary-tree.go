/*
 * @lc app=leetcode id=236 lang=golang
 *
 * [236] Lowest Common Ancestor of a Binary Tree
 */
/**
 * Definition for TreeNode.
 * type TreeNode struct {
 *     Val int
 *     Left *ListNode
 *     Right *ListNode
 * }
 */
 func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
  if root == nil || root.Val == p.Val || root.Val == q.Val{
	  return root
  } 
  left := lowestCommonAncestor(root.Left,p,q)
  right := lowestCommonAncestor(root.Right,p,q)
  // 此处为空的意义？？？
  if left == nil {
	  return right
  }else if right == nil{
	  return left
  }
  return root
}

