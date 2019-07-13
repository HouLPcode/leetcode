/*
 * @lc app=leetcode id=143 lang=golang
 *
 * [143] Reorder List
 */
/**
* Definition for singly-linked list.
* type ListNode struct {
*     Val int
*     Next *ListNode
* }
 */
//  8 ms 97.35 %
func reorderList(head *ListNode) {
	// 快慢指针找到中间节点，然后反转后半部分链表，最后拼接两个链表
	if head == nil || head.Next == nil {
		return // 小于三个节点不需要处理
	}
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
	}
	// if fast == nil {
	// 	// 奇数个节点，slow是中间节点
	// }
	// reverse list
	// 注意，反转前断开链表
	var pre, cur *ListNode = nil, slow.Next
	slow.Next = nil
	for cur != nil {
		pre, cur, cur.Next = cur, cur.Next, pre
	} // pre是链头

	// merge
	curl, curr := head, pre
	for curl != nil && curr != nil {
		curl, curl.Next, curr, curr.Next = curl.Next, curr, curr.Next, curl.Next
	}
	// 没有后续节点要处理
}
