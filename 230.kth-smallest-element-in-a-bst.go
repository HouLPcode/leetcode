/*
 * @lc app=leetcode id=230 lang=golang
 *
 * [230] Kth Smallest Element in a BST
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthSmallest(root *TreeNode, k int) int {
    // 法1 BST中序输出有序序列，选出索要节点，中序遍历采用非递归方式
    // 中序非递归
    if root == nil{
        return 0
    }
    stack := []*TreeNode{}
    // for len(stack) != 0 && root != nil{
    for len(stack) != 0 || root != nil{ //注意此处的条件：或
        //根节点入栈，然后所有左子树入栈
        if root != nil{
            stack = append(stack, root)
            root = root.Left
        } else{
            //出栈的节点即为最终节点
            root = stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            // 访问节点（左节点）
            if k == 1{
                break
            }else{
                k--
            }
            // 指向右子树
            root = root.Right
        }
    }
    return root.Val
}
