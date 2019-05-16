/*
 * @lc app=leetcode id=515 lang=golang
 *
 * [515] Find Largest Value in Each Tree Row
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 失败示例: [0,-1]
// 注意节点的Val有正负，计算最大值时不能初始化为0值
func largestValues(root *TreeNode) []int {
	//按层遍历，统计每行的最大值
	if root == nil {
		return []int{}
	}
	queue := list.New()
	rnt := []int{}
	queue.PushBack(root)
	for queue.Front() != nil{
		l := queue.Len() //统计长度的时候是不是遍历了一遍队列，此处可以继续优化
		// max := 0 // 节点可以有负值，此处不能使用0值
		max := queue.Front().Value.(*TreeNode).Val
		for i:=0; i < l; i++{
			node := queue.Front()
			queue.Remove(node)
			// 注意node.Value才能回去到信息，同时需要格式断言
			if node.Value.(*TreeNode).Val > max{
				max = node.Value.(*TreeNode).Val
			}
			if node.Value.(*TreeNode).Left != nil{
				queue.PushBack(node.Value.(*TreeNode).Left)
			}
			if node.Value.(*TreeNode).Right != nil{
				queue.PushBack(node.Value.(*TreeNode).Right)
			}
		}
		rnt = append(rnt, max)		
	}
	return rnt
}

