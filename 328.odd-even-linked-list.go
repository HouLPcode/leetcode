/*
 * @lc app=leetcode id=328 lang=golang
 *
 * [328] Odd Even Linked List
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
//  (4 ms) √ Your runtime beats 94 %
// 遍历过程中删除偶数节点，构造成偶数链表，然后合并
// 奇数节点是原链表
func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	even := head.Next
	curodd, cureven := head, even

	for cureven != nil && cureven.Next != nil {
		curodd.Next = cureven.Next
		curodd = curodd.Next
		cureven.Next = curodd.Next
		cureven = cureven.Next
	}
	curodd.Next = even
	return head
}
