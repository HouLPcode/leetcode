/*
 * @lc app=leetcode id=235 lang=golang
 *
 * [235] Lowest Common Ancestor of a Binary Search Tree
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
  // 方法1，递归实现
  // 根节点比要找的节点都大，则公共祖先肯定在左子树中
  if root.Val > p.Val && root.Val > q.Val{
	return lowestCommonAncestor(root.Left, p, q)
  }
  // 根节点比要找的节点都大，则公共祖先肯定在左子树中
  if root.Val < p.Val && root.Val < q.Val{
	return lowestCommonAncestor(root.Right, p, q)
  }
  return root 
  // 方法2，非递归实现
}

