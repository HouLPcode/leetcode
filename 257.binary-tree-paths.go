/*
 * @lc app=leetcode id=257 lang=golang
 *
 * [257] Binary Tree Paths
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func binaryTreePaths(root *TreeNode) []string {
  if root == nil{
		return []string{}
	}
	// 不使用辅助函数怎么做？？？？？
	rnt := []string{}
	// 注意传入的指针和string的对应关系
	path(root,fmt.Sprint(root.Val),&rnt)// string(root.Val)可能是\u0001
	// path(root,"",&rnt)
	return rnt
}

// []string需要是指针类型
func path(root *TreeNode, tmp string,rnt *[]string){
	if root.Left == nil && root.Right == nil{
		//  ["->1->2->5","->1->3"]
		*rnt = append(*rnt,tmp)
		return
	}
	if root.Left != nil{
		path(root.Left, tmp+"->"+fmt.Sprint(root.Left.Val), rnt)
	}
	if root.Right != nil{
		path(root.Right, tmp+"->"+fmt.Sprint(root.Right.Val), rnt)
	}
}

