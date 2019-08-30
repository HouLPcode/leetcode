/*
 * @lc app=leetcode id=662 lang=golang
 *
 * [662] Maximum Width of Binary Tree
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
//  52 ms, faster than 12.50% 
// [1]
 // 层次便利
 func widthOfBinaryTree(root *TreeNode) int {
    if root == nil {
		return 0
	}
	queue := []*TreeNode{root}
	max := 1
	for len(queue) > 0 {
		size := len(queue)
		for i:=0; i<size; i++ {
			node := queue[0]
			queue = queue[1:]
			if node == nil {
			    queue = append(queue, nil, nil)
			} else {
				queue = append(queue, node.Left, node.Right)
			}
		}
		//删除首尾的nil
		for len(queue) > 0 && queue[0] == nil {
				queue = queue[1:]
		}
        for len(queue)>0 && queue[len(queue)-1] == nil {
				queue = queue[:len(queue)-1]
		}

		if len(queue) > max {
			max = len(queue)
		}
	}
	return max
}

