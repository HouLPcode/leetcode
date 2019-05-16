/*
 * @lc app=leetcode id=113 lang=golang
 *
 * [113] Path Sum II
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 // 典型的错误，[]int的底层实现机制不熟悉导致的
// 测试集    [5,4,8,11,null,13,4,7,2,null,null,5, 1  ] 22 会出错
// 测试集改为 [5,4,8,11,null,13,4,7,2,null,null,5,null] 22 不回出错
// 测试集改为 [5,4,8,11,null,13,4,2,7,null,null,5,1] 22 也会出错，此时之前对的那条路径也出现了path覆盖的错误
// 对于最终输出结果使用的slice，一定要进行深拷贝，避免后续操作的影响
func pathSum(root *TreeNode, sum int) [][]int {
	if root == nil{
		return [][]int{}
	}
	rnt := [][]int{}
	paths(root,sum,[]int{},&rnt)
	return rnt
}
//path传递的错误
func paths(root *TreeNode,sum int, path []int, rtn *[][]int){
	if root == nil{
		return
	}
	path = append(path,root.Val)
	if root.Val == sum && root.Left == nil && root.Right == nil{
		//注意，此处必须深拷贝，避免之后path的操作对rtn产生影响
		outpath := make([]int, len(path))
		copy(outpath,path)
		(*rtn) = append((*rtn), outpath)//此处的rtn直接饮用的path，所以在其他地方修改path，此处的rtn信息会发生变化，覆盖原先的信息，但是信息不会减少
		return
	}
	paths(root.Left, sum-root.Val, path, rtn)
	paths(root.Right, sum-root.Val, path, rtn) //这个path执行的时候和上一行的还是一样的吗
}
