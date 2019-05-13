/*
 * @lc app=leetcode id=206 lang=golang
 *
 * [206] Reverse Linked List
 *
 * https://leetcode.com/problems/reverse-linked-list/description/
 *
 * algorithms
 * Easy (54.07%)
 * Total Accepted:    564.6K
 * Total Submissions: 1M
 * Testcase Example:  '[1,2,3,4,5]'
 *
 * Reverse a singly linked list.
 * 
 * Example:
 * 
 * 
 * Input: 1->2->3->4->5->NULL
 * Output: 5->4->3->2->1->NULL
 * 
 * 
 * Follow up:
 * 
 * A linked list can be reversed either iteratively or recursively. Could you
 * implement both?
 * 
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

 // 翻转链表，循环和递归实现
 // go 支持多值同时赋值，不需要考虑赋值顺序
 // TODO 递归实现
func reverseList(head *ListNode) *ListNode {
	var cur,prev *ListNode = head,nil // 注意赋值 pre=nil

	for cur != nil {
		cur.Next,prev,cur = prev,cur,cur.Next //同时赋值，顺序不要紧
		// prev,cur,cur.Next = cur,cur.Next,prev
	}
	return prev	
}

// 不支持同时赋值的实现
// public Node reverseList(Node head) {
// 	Node pre = null;
// 	Node next = null;
// 	while (head ! = null) {
// 			next = head.next;
// 			head.next = pre;
// 			pre = head;
// 			head = next;
// 	}
// 	return pre;
// }

