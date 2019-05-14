/*
 * @lc app=leetcode id=114 lang=golang
 *
 * [114] Flatten Binary Tree to Linked List
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func flatten(root *TreeNode)  {
    //前序遍历
    if root == nil{
        return
    }
    stack := []*TreeNode{}
    stack = append(stack,root)
    // head := 
    p := &TreeNode{} //空头节点
    for len(stack) > 0{
        node := stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        p.Right = node
        p.Left = nil
        p = p.Right
        
        if node.Right != nil{
            stack = append(stack,node.Right)
        }
        if node.Left != nil{
            stack = append(stack,node.Left)
        }
    }
    // 前序root肯定是头节点，创建的空头节点可以不处理
}

