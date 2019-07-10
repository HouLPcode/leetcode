/*
 * @lc app=leetcode id=24 lang=golang
 *
 * [24] Swap Nodes in Pairs
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
//  0 ms
// recursion 递归实现
func swapPairs(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    // step1: swap
    head, head.Next.Next, head.Next  = head.Next, head, head.Next.Next
    // step2: recursion
    head.Next.Next = swapPairs(head.Next.Next)
    return head
}

