/*
 * @lc app=leetcode id=98 lang=golang
 *
 * [98] Validate Binary Search Tree
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
    //注意，不能只看左节点<跟<右节点，需要时左子树<根<右子树
	// 方法1：中序遍历，判断是否有序，（优化：只保存1个前继节点即可）
	
	// 方法2：递归实现
	// TODO int类型的最大值和最小值，此处不知道int占几字节
	// 所以不考虑math.MaxInt32这种形式
	const MaxInt int = (1 << (bits.UintSize-1)) - 1
	const MinInt int = -(1 << (bits.UintSize-1))

	return valid(root,MinInt,MaxInt)
}

func valid(root *TreeNode,min,max int)bool{
	if root == nil{
		return true
	}
	if root.Val <= min || root.Val >= max {
		return false
	}
	return valid(root.Left, min, root.Val) && valid(root.Right, root.Val, max)
}

