/*
 * @lc app=leetcode id=1110 lang=golang
 *
 * [1110] Delete Nodes And Return Forest
 */

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 12 ms, faster than 95.87%
// [1,2,4,null,3]\n[3]
// var ans = []*TreeNode{} // 全局变量作为最终的返回
func delNodes(root *TreeNode, to_delete []int) []*TreeNode {
	dels := make(map[int]bool, len(to_delete))
	for _, v := range to_delete {
		dels[v] = true
	}
	ans := []*TreeNode{}

	// del(root, dels) 典型错误，这样会少算当前树的根节点
	// return ans
	root = del(root, dels, &ans)
	if root != nil {
		ans = append(ans, root)
	}

	return ans
}

func del(root *TreeNode, dels map[int]bool, ans *([]*TreeNode)) *TreeNode {
	if root == nil {
		return root
	}
	// left,right,root 后序
	root.Left = del(root.Left, dels, ans)
	root.Right = del(root.Right, dels, ans)
	if dels[root.Val] == true {
		if root.Left != nil {
			*ans = append(*ans, root.Left)
		}
		if root.Right != nil {
			*ans = append(*ans, root.Right)
		}
		return nil
	}
	return root
}
