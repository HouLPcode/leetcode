/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// err--------------------
// 典型错误，思路错误，不是隔一层一个摄像头，是两层
// [0]
// [0,0,null,null,0,0,null,null,0,0] 是2个
// 树状DP
func minCameraCover(root *TreeNode) int {
	if root == nil {
		return 0
	}
	// root 安装
	cnt0 := place(root)
	// root不安装
	if root.Left != nil || root.Right != nil {
		// 叶子节点只能安装
		cnt1 := place(root.Left) + place(root.Right)
		return min(cnt0, cnt1)
	}
	return cnt0
}

// 在root安装需要的个数
func place(root *TreeNode) int {
	if root == nil {
		return 0
	}

	// root 安装摄像头
	cnt := 1
	if root.Left != nil {
		cnt += place(root.Left.Left)
		cnt += place(root.Left.Right)
	}
	if root.Right != nil {
		cnt += place(root.Right.Left)
		cnt += place(root.Right.Right)
	}
	return cnt
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
