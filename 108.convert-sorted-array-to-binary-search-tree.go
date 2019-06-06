/*
 * @lc app=leetcode id=108 lang=golang
 *
 * [108] Convert Sorted Array to Binary Search Tree
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedArrayToBST(nums []int) *TreeNode {
	// 折半递归创建
	if len(nums) == 0{
		return nil
	}
	// 怎么保证创建的是平衡树
	// AVL树的调整设计到左右旋转问题
	// 因为常见节点的时候是尽可能的左右节点都有，创建的就是平衡树
	return &TreeNode{
		Val: nums[len(nums)/2],
		Left: sortedArrayToBST(nums[:len(nums)/2]),
		Right: sortedArrayToBST(nums[len(nums)/2+1:]),// +1不会越界
	}
}

