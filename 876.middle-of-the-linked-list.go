/*
 * @lc app=leetcode id=876 lang=golang
 *
 * [876] Middle of the Linked List
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
//  0 ms
func middleNode(head *ListNode) *ListNode {
	// 快慢指针
	s1, s2 := head, head.Next // 题中要求head非空,注意初始值的设置
	for ; s2 != nil; s1, s2 = s1.Next, s2.Next.Next{
		if s2.Next == nil {
			break
		}
	}
	if s2 == nil {
		return s1
	} 
	return s1.Next
}

