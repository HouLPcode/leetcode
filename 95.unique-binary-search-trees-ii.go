/*
 * @lc app=leetcode id=95 lang=golang
 *
 * [95] Unique Binary Search Trees II
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
//28 ms, faster than 38.55%
// 0
func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return []*TreeNode{}
	}
	return bst(1, n)
}

func bst(s, e int) []*TreeNode {
	// 暂时先不用缓存加速
	if s > e {
		return []*TreeNode{nil} // 注意，是nil而不是空
	}
	// root = [s,e]
	rtn := []*TreeNode{}
	for r := s; r <= e; r++ {
		for _, ltree := range bst(s, r-1) {
			for _, rtree := range bst(r+1, e) {
				root := &TreeNode{r, ltree, rtree}
				rtn = append(rtn, root)
			}
		}
	}
	return rtn
}
