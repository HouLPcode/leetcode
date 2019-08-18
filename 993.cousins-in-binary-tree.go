/*
 * @lc app=leetcode id=993 lang=golang
 *
 * [993] Cousins in Binary Tree
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// (0 ms) √ Your runtime beats 100 %
// BFS 层次遍历。
// 在子节点中找到 x，y值，则将该节点存储在parent中，一层遍历完之后，根据parent的情况返回结果
func isCousins(root *TreeNode, x int, y int) bool {
	// if root == nil{
	//     return
	// }

	queue := []*TreeNode{}      //创建队列
	queue = append(queue, root) // 入队
	parent := make([]int, 0, 2) // 记录每个节点的parent
	for len(queue) != 0 {
		l := len(queue)
		for i := 0; i < l; i++ {
			node := queue[0] //出队
			queue = queue[1:]

			//子节点入队
			if node.Left != nil {
				queue = append(queue, node.Left)
				if node.Left.Val == x || node.Left.Val == y {
					parent = append(parent, node.Val)
				}
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
				if node.Right.Val == x || node.Right.Val == y {
					parent = append(parent, node.Val)
				}
			}
		}
		// 遍历完一层节点后
		if len(parent) == 2 && parent[0] != parent[1] {
			return true
		} else if len(parent) != 0 {
			return false
		}
	}
	return false
}
