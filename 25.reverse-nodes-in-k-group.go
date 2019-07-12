/*
 * @lc app=leetcode id=25 lang=golang
 *
 * [25] Reverse Nodes in k-Group
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
//  8 ms  23.97 % 递归实现
func reverseKGroup(head *ListNode, k int) *ListNode {
	if k < 2 {
		return head
	}

	i := 0
	// 查看节点个数够不够k个
	for p := head; i < k && p != nil; p = p.Next {
		i++
	}
	if i < k { // 一定注意，此处不能通过 p == nil判断，2个节点，k=2，则p也指向nil，导致末尾未反转
		return head // 不够
	}
	// reverse k nodes
	var pre, cur *ListNode = nil, head
	for i := 0; i < k; i++ {
		pre, cur, cur.Next = cur, cur.Next, pre
	}
	head.Next = reverseKGroup(cur, k)
	return pre
}
