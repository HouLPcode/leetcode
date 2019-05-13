/*
 * @lc app=leetcode id=61 lang=golang
 *
 * [61] Rotate List
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
	//一次遍历，统计节点个数并构成环链表，k%nodenum减少循环次数
	// nil 处理
	if head == nil{
		return nil
	}
	nodeNum := 1// 非空至少有1个节点
	p := head
	for ; p.Next != nil; p = p.Next{
		nodeNum++
	}
	p.Next = head // 构成环

	k = nodeNum - k % nodeNum // 正向遍历的个数
	for k > 1{ // 此处获取到新头结点的前1个节点
		head = head.Next
		k--
	}
	head, head.Next = head.Next, nil // 此处必须同时赋值
	return head
}

