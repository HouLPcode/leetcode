/*
 * @lc app=leetcode id=107 lang=golang
 *
 * [107] Binary Tree Level Order Traversal II
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrderBottom(root *TreeNode) [][]int {
	// 递归实现，新的level时，在前面插入
	rnt := [][]int{}
	levelOrder(root, 1, &rnt)
	return rnt
}

func levelOrder(root *TreeNode, level int, rnt *[][]int){
	if root == nil{
		return 
	}
	if len((*rnt)) < level{//新的一行，插入最前面
		(*rnt) = append([][]int{{root.Val},},(*rnt)... )
	}else{// 添加到 len-level中
		(*rnt)[len((*rnt))-level] = append((*rnt)[len((*rnt))-level],root.Val)
	}
	levelOrder(root.Left, level+1, rnt)
	levelOrder(root.Right, level+1, rnt)
}

