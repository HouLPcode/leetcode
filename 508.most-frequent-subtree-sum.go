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
	tmp := make(map[int]int)
	sum(root, tmp)
	maxTime := -1
	rtn := []int{}
	for k, v := range tmp {
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
	l, r := 0, 0
	if root.Left != nil {
		l = sum(root.Left, tmp)
	}
	if root.Right != nil {
		r = sum(root.Right, tmp)
	}

	tmp[l+r+root.Val]++ // 只在此处统计一次结果，在上面的left或right判断中不记录，否则会导致多计算一次
	return l + r + root.Val
}
