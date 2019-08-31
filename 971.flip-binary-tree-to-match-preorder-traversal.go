/*
 * @lc app=leetcode id=971 lang=golang
 *
 * [971] Flip Binary Tree To Match Preorder Traversal
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 0 ms
// [1,2,null,3]\n[1,3,2]
// 通过全局index， 避免区分具体的左右子树节点，每次判断root值后，index++
func flipMatchVoyage(root *TreeNode, voyage []int) []int {
	index := 0
	// 此处通过闭包，访问变量index，这样就不用具体的区分左右子树的数据了
	var dfs func(node *TreeNode) []int
	dfs = func(node *TreeNode) []int {
		if node == nil {
			return []int{}
		}
		if node.Val != voyage[index] {
			return []int{-1}
		}

		res := []int{}
		index++ // 共享index
		if node.Left != nil && node.Left.Val != voyage[index] {
			res = append(res, node.Val)
			flip := dfs(node.Right)
			if len(flip) > 0 && flip[0] == -1 {
				return []int{-1}
			} else {
				res = append(res, flip...)
			}
			flip = dfs(node.Left)
			if len(flip) > 0 && flip[0] == -1 {
				return []int{-1}
			} else {
				res = append(res, flip...)
			}
		} else {
			flip := dfs(node.Left)
			if len(flip) > 0 && flip[0] == -1 {
				return []int{-1}
			} else {
				res = append(res, flip...)
			}
			flip = dfs(node.Right)
			if len(flip) > 0 && flip[0] == -1 {
				return []int{-1}
			} else {
				res = append(res, flip...)
			}
		}
		return res
	}

	return dfs(root)
}
