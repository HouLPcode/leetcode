/*
 * @lc app=leetcode id=203 lang=golang
 *
 * [203] Remove Linked List Elements
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
	// 注意头结点的删除
	for head != nil && head.Val == val{
		head = head.Next
	} // 情况1，所有节点都删除，head可能为nil
	  // 情况2，head原本就是nil
	if head == nil{
		return nil
	}

	p := head
	for p.Next != nil{
		if p.Next.Val == val{
			p.Next = p.Next.Next
		}else{
			p = p.Next
		}
	}
	return head
}

