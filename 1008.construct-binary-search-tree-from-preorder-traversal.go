/*
 * @lc app=leetcode id=1008 lang=golang
 *
 * [1008] Construct Binary Search Tree from Preorder Traversal
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
//  (0 ms)
func bstFromPreorder(preorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	index := 1
	for index < len(preorder) && preorder[index] < preorder[0] {
		index++
	}

	rtn := &TreeNode{
		Val:   preorder[0],
		Left:  bstFromPreorder(preorder[1:index]),
		Right: bstFromPreorder(preorder[index:]),
	}
	return rtn
	// if index == 1 { // 不需要下面的判断
	//     rtn.Right = bstFromPreorder(preorder[index:])
	// } else if index == len(preorder) {
	//     rtn.Left = bstFromPreorder(preorder[1:index])
	// } else {
	//     rtn.Left = bstFromPreorder(preorder[1:index])
	//     rtn.Right = bstFromPreorder(preorder[index:])
	// }

}
