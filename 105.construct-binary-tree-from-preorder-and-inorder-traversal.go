/*
 * @lc app=leetcode id=105 lang=golang
 *
 * [105] Construct Binary Tree from Preorder and Inorder Traversal
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
    if len(preorder) == 0 && len(inorder) == 0{
		return nil
	}
	// 注意 1个节点没必要单独出来
	// if len(preorder) == 1 && len(inorder) == 1{
	// 	return &TreeNode{
	// 		Val:preorder[0],
	// 		Left:nil,
	// 		Right:nil,
	// 	}
	// }

	rootVal := preorder[0]	
	// 没有重复元素
	i := 0
	for ; i < len(inorder); i++{
		if inorder[i] == rootVal{
			break
		}
	}

	// 一个底层slice的多个值赋值会导致什么问题吗？？？	
	var lpreorder,rpreorder,linorder,rinorder []int
	if i == 0{//无左子树
		rpreorder = preorder[1:]
		rinorder = inorder[i+1:] //这样访问不会越界,i+2会越界
	}else if i == len(inorder) -1{//无右子树
		lpreorder = preorder[1:]
		linorder = inorder[:i]
	}else{
		linorder = inorder[:i]
		rinorder = inorder[i+1:]
		rpreorder = preorder[len(linorder)+1:]
		lpreorder = preorder[1:len(linorder)+1]
	}
	
	return &TreeNode{
		Val:rootVal,
		Left:buildTree(lpreorder,linorder),
		Right:buildTree(rpreorder,rinorder)}
}

