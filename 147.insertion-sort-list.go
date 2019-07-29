import "math"

/*
 * @lc app=leetcode id=147 lang=golang
 *
 * [147] Insertion Sort List
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
//  28 ms 55.33%
func insertionSortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	out := &ListNode{math.MinInt64, nil} // 头指针
	for head != nil {
		// 从左往右找到要插入的位置
		cur := out
		for cur.Next != nil && head.Val > cur.Next.Val {
			cur = cur.Next
		}
		cur.Next, head.Next, head = head, cur.Next, head.Next
	}

	return out.Next
}
