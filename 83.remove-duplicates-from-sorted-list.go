/*
 * @lc app=leetcode id=83 lang=golang
 *
 * [83] Remove Duplicates from Sorted List
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil{
		return nil
	}
	p := head
	for p.Next != nil{
		if p.Val == p.Next.Val{// 注意，删除重复元素后p不推进，用来处理连续多个重复值节点
			p.Next = p.Next.Next
		}else{
			p = p.Next
		}
	}
	return head
}

