/*
 * @lc app=leetcode id=113 lang=golang
 *
 * [113] Path Sum II
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, sum int) [][]int {
	rnt := [][]int{}
	path(root,sum,0,[]int{},&rnt)
	return rnt
}

func path(root *TreeNode,sum int, curSum int, tmp []int, rnt *[][]int){
	if root == nil{
		return
	}
	if root.Left == nil && root.Right == nil {
		if curSum+root.Val == sum {
			*rnt = append(*rnt, append(tmp,root.Val))
		}
		return
	}
	if root.Left !=nil{
		path(root.Left,sum,curSum+root.Val,append(tmp,root.Val),rnt)
	}
	if root.Right != nil{
		path(root.Right,sum,curSum+root.Val,append(tmp,root.Val),rnt)
	}
}

