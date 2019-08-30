/*
 * @lc app=leetcode id=655 lang=golang
 *
 * [655] Print Binary Tree
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 // 0ms
 func printTree(root *TreeNode) [][]string {
	rows := depth(root)
	cols :=  int(math.Pow(float64(2), float64(rows))) - 1
	ans := make([][]string, rows)
	for r:=0; r<rows;r++{
		ans[r] = make([]string, cols)
	}
	printNode(root, 0, 0,cols-1, ans)
	return ans
}

// 注意根节点在第一层中间，左子树在左边剩余位置的中间，右子树在剩余位置的中间
// s,e 当前行的起止列
func printNode(root *TreeNode, row, s,e int, ans[][]string) {
	// preOrder
    mid := (s + e)/2
    ans[row][mid] = strconv.Itoa(root.Val)
	if root.Left != nil{
        printNode(root.Left, row+1,s,mid-1, ans)
	}
	if root.Right != nil{
        printNode(root.Right, row+1,mid+1,e, ans)
	}
}


func depth(root *TreeNode) int {
    if root == nil {
        return 0
    }
    return max(depth(root.Left), depth(root.Right))+1
}

func max(a, b int) int {
	if a > b {
		return a 
	}
	return b
}

 

