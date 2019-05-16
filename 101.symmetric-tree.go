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
// 递归实现, 退出条件 a==b==nil, 递归过程中a!=b||a.Val!=b.Val返回false
// 非递归，
    // 左->入队，右 <- 入队，每次出队两个元素，判断是否一致，不用分层实现
func isSymmetric(root *TreeNode) bool {

    // root==nil的判断下放到isMirror函数中
    return isMirror(root, root)//如果root是对称的，则自己和自己也是镜像的
}
func isMirror(lTree,rTree *TreeNode) bool{
    if lTree==nil && rTree==nil{
        return true
    }
    queue := list.New()
    queue.PushBack(lTree)
    queue.PushBack(rTree)
    for queue.Front() != nil {
        lnode := queue.Front()
        queue.Remove(lnode)
        rnode := queue.Front()
        queue.Remove(rnode)
        if (lnode.Value.(*TreeNode).Val != rnode.Value.(*TreeNode).Val) ||
            (lnode.Value.(*TreeNode).Left == nil && rnode.Value.(*TreeNode).Right != nil) ||
            (lnode.Value.(*TreeNode).Left != nil && rnode.Value.(*TreeNode).Right == nil){
            return false         //注意指针是不是nil对称的比较，不能只看Val
        }
        if lnode.Value.(*TreeNode).Left != nil && rnode.Value.(*TreeNode).Right != nil{
            queue.PushBack(lnode.Value.(*TreeNode).Left)
            queue.PushBack(rnode.Value.(*TreeNode).Right)
        }
        if lnode.Value.(*TreeNode).Right != nil && rnode.Value.(*TreeNode).Left != nil{
            queue.PushBack(lnode.Value.(*TreeNode).Right)
            queue.PushBack(rnode.Value.(*TreeNode).Left)
        }
    }
    return true
}
