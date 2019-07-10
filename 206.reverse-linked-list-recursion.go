/*
 * @lc app=leetcode id=206 lang=golang
 *
 * [206] Reverse Linked List
 *
 */
//  0 ms
 // 递归实现
func reverseList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil{
        return head
    }
    // 同时赋值，保证连续性，避免前后顺序导致出错
    head, head.Next, head.Next.Next = reverseList(head.Next), nil, head
    return head
}

