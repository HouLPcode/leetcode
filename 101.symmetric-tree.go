/*
 * @lc app=leetcode id=101 lang=golang
 *
 * [101] Symmetric Tree
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
    //递归实现
    // root==nil的判断下放到isMirror函数中
    return isMirror(root, root)//如果root是对称的，则自己和自己也是镜像的
}
func isMirror(a,b *TreeNode) bool{
    if a==nil && b==nil{
        return true
    }
    if a == nil || b == nil{//经过前面的判断，此处一个空一个非空
        return false
    }
    // a，b都不是nil
    return a.Val == b.Val && isMirror(a.Left, b.Right) && isMirror(a.Right, b.Left)
}
