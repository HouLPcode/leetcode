/*
 * @lc app=leetcode id=654 lang=golang
 *
 * [654] Maximum Binary Tree
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
//  (56 ms) √ Your runtime beats 77.92 %
func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	max := nums[0]
	maxIndex := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] > max {
			max = nums[i]
			maxIndex = i
		}
	}
	rtn := &TreeNode{
		Val:   max,
		Left:  nil,
		Right: nil}
	if maxIndex > 0 && maxIndex < len(nums) { // 注意不饿能忘了此处的判断，否则切片可能导致访问越界
		rtn.Left = constructMaximumBinaryTree(nums[:maxIndex])
	}
	if maxIndex <= len(nums)-2 && maxIndex >= 0 {
		rtn.Right = constructMaximumBinaryTree(nums[maxIndex+1:])
	}
	return rtn
}
