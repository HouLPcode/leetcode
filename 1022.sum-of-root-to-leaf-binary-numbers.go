/*
 * @lc app=leetcode id=1022 lang=golang
 *
 * [1022] Sum of Root To Leaf Binary Numbers
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
//  0 ms 
// 错误思路：层次遍历，统计每层中1的个数，然后二进制计算总和。路径深度不同，导致计算错误
// 先获取到所有的paths，然后加和
func sumRootToLeaf(root *TreeNode) int {
	paths := [][]int{}
	dfs(root, []int{}, &paths)
	return calcPaths(paths)
}

func dfs(root *TreeNode, cur []int, paths*[][]int) {
	if root == nil {
		return 
	}
	cur = append(cur, root.Val)
	if root.Left == nil && root.Right == nil {
		tmp := make([]int, len(cur))
		copy(tmp, cur)
		(*paths) = append((*paths), tmp)
		return 
	}
	if root.Left != nil {
		dfs(root.Left, cur, paths)
	}
	if root.Right != nil {
		dfs(root.Right, cur, paths)
	}
	cur = cur[:len(cur)-1]
}

func calcPaths(paths [][]int) int {
	rtn := 0 
	for _,v := range paths {
		rtn += bit2int(v)
	}
	return rtn
}

func bit2int(data []int) int {
	rtn := 0
	for _,v := range data {
		rtn = rtn << 1 + v
	}
	return rtn
}

