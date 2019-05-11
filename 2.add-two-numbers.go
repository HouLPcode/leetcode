/*
 * @lc app=leetcode id=2 lang=golang
 *
 * [2] Add Two Numbers
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	//list ->> int 有可能越界，不考虑
	// 结果直接存储在l1链表中
	c := int(0) // 进位
	head := l1

	for l1 != nil && l2 !=nil{
		l1.Val, c = (l1.Val + l2.Val + c) %10, (l1.Val + l2.Val + c) /10 //必须同时赋值，计算c的时候l1.Val已经改变
		l1, l2 = l1.Next, l2.Next
	}

	// 两个链表长度不同
	for l1 != nil{
		l1.Val, c = (l1.Val + c) % 10, (l1.Val + c) / 10
		l1 = l1.Next
	}
	if l2 != nil{
		p := head
		for ; p.Next!=nil;p = p.Next{}
		p.Next = l2
	}
	for l2 != nil{
		l2.Val, c = (l2.Val + c) % 10, (l2.Val + c) / 10
		l2 = l2.Next
	}

	 // 注意最后的进位，需要添加一个新的节点
	if c != int(0){
		p := head
		for ; p.Next!=nil;p = p.Next{}
		p.Next = &ListNode{c,nil}
	}
	return head
}

