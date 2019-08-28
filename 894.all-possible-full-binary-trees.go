/*
 * @lc app=leetcode id=894 lang=golang
 *
 * [894] All Possible Full Binary Trees
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
//  432 ms, faster than 70.00%
// 满二叉树的总结点个数规律 , 奇数节点
func allPossibleFBT(N int) []*TreeNode {
	// map 加速
	m := make(map[int][]*TreeNode)
	return FBT(N, m)
}

func FBT(N int, m map[int][]*TreeNode) []*TreeNode {
	if v, ok := m[N]; ok {
		return v
	}
	rtn := []*TreeNode{}
	if N == 1 {
		rtn = append(rtn, &TreeNode{
			Val:   0,
			Left:  nil,
			Right: nil,
		})
	} else if N&1 != 0 {
		for x := 0; x < N; x++ {
			y := N - 1 - x
			tree := &TreeNode{}
			tree.Val = 0
			for _, ltree := range FBT(x, m) {
				for _, rtree := range FBT(y, m) {
					rtn = append(rtn, &TreeNode{
						Val:   0,
						Left:  ltree,
						Right: rtree,
					})
				}
			}
		}
	}
	m[N] = rtn
	return rtn
}
