/*
 * @lc app=leetcode id=99 lang=golang
 *
 * [99] Recover Binary Search Tree
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// err------------------------------
// [0,1]
// [3,1,2]
// [5,3,9,-2147483648,2]
// 递归，中序缓存所有节点，然后找到需要交换的节点
func recoverTree(root *TreeNode) {
	nodes := []*TreeNode{}
	inOrder(root, &nodes)

	if len(nodes) == 2 {
		nodes[0].Val, nodes[1].Val = nodes[1].Val, nodes[0].Val
		return
	}

	// 注意怎么找两个需要减缓的节点的
	// first, 找到a>b处，选择a
	// secend, 继续找 c>d, 选择d， 如果只剩一个数，就选则这个数
	first := 0
	for nodes[first].Val <= nodes[first+1].Val {
		first++
	}
	secend := first + 1
	if secend != len(nodes)-1 {
		for secend < len(nodes)-1 && nodes[secend].Val <= nodes[secend+1].Val {
			secend++
		}
		if secend == len(nodes)-1 {
			secend-- //遍历到末尾，则选择最后一个数进行交换
		} else {
			secend++ // 取最后一个
		}
	}

	nodes[first].Val, nodes[secend].Val = nodes[secend].Val, nodes[first].Val
}

func inOrder(root *TreeNode, ans *[]*TreeNode) {
	if root == nil {
		return
	}
	inOrder(root.Left, ans)
	*ans = append(*ans, root)
	inOrder(root.Right, ans)
}
