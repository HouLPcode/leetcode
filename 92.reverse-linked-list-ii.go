/*
 * @lc app=leetcode id=92 lang=golang
 *
 * [92] Reverse Linked List II
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, m int, n int) *ListNode {
	var pri, cur *ListNode = nil, head 	// pri指向m前一个节点，cur指向m节点
	for i:=0; i< m-1; i++{//此处不要直接m--，会影响后面的使用
		pri,cur = cur, cur.Next
	}

	pripri := pri //pripri指向翻转节点的前一个节点，可以是nil
	for cnt:=n-m+1;cnt>0;cnt--{
		pri,cur,cur.Next = cur,cur.Next,pri
	}
	
	if pripri == nil{//头节点参与翻转
		head,head.Next = pri, cur
	}else{
		pripri.Next, pripri.Next.Next = pri, cur
	}
	return head 
}

