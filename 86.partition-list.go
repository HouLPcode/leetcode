/*
 * @lc app=leetcode id=86 lang=golang
 *
 * [86] Partition List
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {
	if head == nil{
		return nil
	}
	var p1,p2 *ListNode = &ListNode{}, &ListNode{} // 空头指针
	head1 := p1
	head2 := p2
	for p:=head; p!=nil; p = p.Next{
		if p.Val < x{
			p1.Next = p
			p1 = p1.Next
		}else{
			p2.Next = p
			p2 = p2.Next
		}
	}
	p2.Next = nil
	p1.Next = head2.Next //去掉p2头节点
	return head1.Next
}

