/*
 * @lc app=leetcode id=106 lang=golang
 *
 * [106] Construct Binary Tree from Inorder and Postorder Traversal
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(inorder []int, postorder []int) *TreeNode {
    if len(inorder) == 0 && len(postorder) == 0{
		// return &TreeNode{}
		return nil //返回nil而不是空节点，空节点{0,nil,nil}
	}

	rootVal := postorder[len(postorder)-1]

	//在inorder中找到左右子树
	i:=0
	for ;i<len(inorder);i++{
		if inorder[i] == rootVal{
			break
		}
	}
	 
	var linorder,rinorder,lpostorder,rpostorder []int
	if i == 0 {//无左   单个节点无左无右
		rinorder = inorder[i+1:]//单节点也不会越界
		rpostorder = postorder[:len(postorder)-1]
	}else if i == len(inorder) - 1 {//无右
		linorder = inorder[:i]
		lpostorder =  postorder[:i]
	}else{
		linorder = inorder[:i]
		rinorder = inorder[i+1:]
		lpostorder = postorder[:i]
		rpostorder = postorder[i:len(postorder)-1]
	}
	return &TreeNode{
		Val : rootVal,
		Left : buildTree(linorder, lpostorder),
		Right : buildTree(rinorder, rpostorder),
	}
}

