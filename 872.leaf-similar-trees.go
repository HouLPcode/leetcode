/*
 * @lc app=leetcode id=872 lang=golang
 *
 * [872] Leaf-Similar Trees
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 0 ms

//  [1]\n[2]
func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	// 中序遍历得到两树的叶子序列，然后比较
	nodes1 := make([]int, 0)
	nodes2 := make([]int, 0)
	inOrder(root1, &nodes1)
	inOrder(root2, &nodes2)
	return reflect.DeepEqual(nodes1, nodes2)
}

// nodes 通过[]int传递会出现错误
func inOrder(root *TreeNode, nodes *[]int) {
	if root== nil {
		return 
	}
	if root.Left != nil {
		inOrder(root.Left, nodes)
	}
	if root.Left == nil && root.Right == nil {
		(*nodes) = append((*nodes), root.Val)
	}
	if root.Right != nil {
		inOrder(root.Right, nodes)
	}
}

