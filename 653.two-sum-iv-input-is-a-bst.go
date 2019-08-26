/*
 * @lc app=leetcode id=653 lang=golang
 *
 * [653] Two Sum IV - Input is a BST
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 20ms 96.92%
// 数组缓存中序结果，然后两端往中间查找，遍历结果
func findTarget(root *TreeNode, k int) bool {
	if root == nil {
		return false
	}
	tmp := inOrder(root)
	s, e := 0, len(tmp)-1
	for s >= 0 && e < len(tmp) && s < e {
		if tmp[s]+tmp[e] > k {
			e--
		} else if tmp[s]+tmp[e] == k {
			return true
		} else {
			s++
		}
	}
	return false
}

// 非递归
func inOrder(root *TreeNode) []int {
	rtn := []int{}
	if root == nil {
		return []int{}
	}
	stack := []*TreeNode{}              // 初始空栈，不需要root入栈
	for len(stack) > 0 || root != nil { // 注意此处不能只判断栈空
		if root != nil { // 注意此处不能是循环
			stack = append(stack, root)
			root = root.Left
		} else {
			// 出栈
			root, stack = stack[len(stack)-1], stack[:len(stack)-1]
			// 访问根节点
			rtn = append(rtn, root.Val)
			// 指向右节点
			root = root.Right
		}
	}
	return rtn
}
