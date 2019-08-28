/*
 * @lc app=leetcode id=889 lang=golang
 *
 * [889] Construct Binary Tree from Preorder and Postorder Traversal
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 4 ms, faster than 85.29%
func constructFromPrePost(pre []int, post []int) *TreeNode {
	if len(pre) == 0 {
		return nil
	}
	// 前序，后序构造树，返回任意一个都行
	rtn := &TreeNode{
		Val: pre[0],
	}
	if len(pre) == 1 {
		return rtn
	}

	// i := 1
	// lmap, rmap := make(map[int]bool), make(map[int]bool)
	// for i < len(pre) {
	// 	lmap[pre[i]] = true
	// 	rmap[post[i-1]] = true
	// 	if reflect.DeepEqual(lmap, rmap) { // 待优化
	// 		break
	// 	}
	// 	i++
	// }

	L := 1
	for ; L < len(pre); L++ {
		if post[L-1] == pre[1] {
			break
		}
	}

	rtn.Left = constructFromPrePost(pre[1:L+1], post[0:L])
	rtn.Right = constructFromPrePost(pre[L+1:], post[L:])

	return rtn
}
