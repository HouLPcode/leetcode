/*
 * @lc app=leetcode id=237 lang=golang
 *
 * [237] Delete Node in a Linked List
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 // 题目中给的node是需要删除的node，没有给出head，不需要遍历
func deleteNode(node *ListNode) {
	// node 不是最后一个
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}

