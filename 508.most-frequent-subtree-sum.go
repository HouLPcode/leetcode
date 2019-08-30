/*
 * @lc app=leetcode id=508 lang=golang
 *
 * [508] Most Frequent Subtree Sum
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 12 ms, faster than 68.85%
// map缓存所有结果
func findFrequentTreeSum(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	counter := make(map[int]int)
	sum(root, counter)
	maxTime := -1
	rtn := []int{}
	for k, v := range counter {
		if v > maxTime {
			maxTime = v
			rtn = append(rtn, k)
			rtn = rtn[len(rtn)-1:]
		} else if v == maxTime {
			rtn = append(rtn, k)
		}
	}
	return rtn
}

func sum(root *TreeNode, tmp map[int]int) int {
	if root == nil {
		return 0
	}
	data := sum(root.Left, tmp) + sum(root.Right, tmp) + root.Val
	tmp[data]++
	return data
}
