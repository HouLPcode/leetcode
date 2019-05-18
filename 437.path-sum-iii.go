/*
 * @lc app=leetcode id=437 lang=golang
 *
 * [437] Path Sum III
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 //方法1，dfs，记录所有路径，然后路径中找path，最后计数。数组中找sum题目中用到了map缓存
 //方法2，dfs，递归的时候分要和不要当前节点， 
func pathSum(root *TreeNode, sum int) int {
	//方法2 尝试
	target = sum
	count := 0
	path(root, target, &count)
	return count
}
var target int //全局存储题目中的sum，保持不变

func path(root *TreeNode, sum int, count *int){
	if root == nil{
		return 
	}
	if root.Val == sum{
		*count++//???????? * 和 ++
	}// 注意，即便当前节点是满足path的，也要继续向下寻找，有可能后面的节点和为0
	path(root.Left, sum-root.Val, count)//要当前节点
	path(root.Right, sum-root.Val, count)
	path(root.Left, target, count)//不要当前节点，从题中要求的sum重新开始
	path(root.Right, target, count)
}
