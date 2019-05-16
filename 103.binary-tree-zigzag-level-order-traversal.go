/*
 * @lc app=leetcode id=103 lang=golang
 *
 * [103] Binary Tree Zigzag Level Order Traversal
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 典型错误，不能一次遍历从左往右，下一次从右向左，因为出队的序列决定过了子节点在整层遍历中的序列
// 方法1， 正常增次遍历，之后隔行将[]int翻转
// 方法2， 双端链表替代普通queue，每次从不同的方向遍历
		// 关注出队顺序，左右，右左，
		// 入队顺序，左右，右左，前面入队还是后面入队，
func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil{
		return [][]int{}
	}

	rnt := [][]int{}
	queue := list.New()
	queue.PushBack(root)
	for queue.Front() != nil{
		tmp := []int{}
		l := queue.Len()
		for i:=0; i<l; i++{
			node := queue.Front() //左右出队
			queue.Remove(node)
			tmp = append(tmp, node.Value.(*TreeNode).Val)
			if node.Value.(*TreeNode).Left != nil{ //左右后面入队
				queue.PushBack(node.Value.(*TreeNode).Left)
			}
			if node.Value.(*TreeNode).Right != nil{
				queue.PushBack(node.Value.(*TreeNode).Right)
			}	 
		}
		rnt = append(rnt, tmp)
		tmp = []int{}
		l = queue.Len()
		if l == 0{
			break //奇数层 ，退出循环，防止后面对个[]
		}
		for i:=0; i<l; i++{
			node := queue.Back() // 右左出队
			queue.Remove(node) 

			tmp = append(tmp, node.Value.(*TreeNode).Val)

			if node.Value.(*TreeNode).Right != nil{ //右左前面入队
				queue.PushFront(node.Value.(*TreeNode).Right)
			} 
			if node.Value.(*TreeNode).Left != nil{
				queue.PushFront(node.Value.(*TreeNode).Left)
			}  
		}
		rnt = append(rnt, tmp)
	}
	return rnt
}

