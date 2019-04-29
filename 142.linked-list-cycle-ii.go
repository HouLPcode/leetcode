/*
 * @lc app=leetcode id=142 lang=golang
 *
 * [142] Linked List Cycle II
 *
 * https://leetcode.com/problems/linked-list-cycle-ii/description/
 *
 * algorithms
 * Medium (31.69%)
 * Total Accepted:    208.4K
 * Total Submissions: 657.6K
 * Testcase Example:  '[3,2,0,-4]\n1'
 *
 * Given a linked list, return the node where the cycle begins. If there is no
 * cycle, return null.
 * 
 * To represent a cycle in the given linked list, we use an integer pos which
 * represents the position (0-indexed) in the linked list where tail connects
 * to. If pos is -1, then there is no cycle in the linked list.
 * 
 * Note: Do not modify the linked list.
 * 
 * 
 * 
 * Example 1:
 * 
 * 
 * Input: head = [3,2,0,-4], pos = 1
 * Output: tail connects to node index 1
 * Explanation: There is a cycle in the linked list, where tail connects to the
 * second node.
 * 
 * 
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: head = [1,2], pos = 0
 * Output: tail connects to node index 0
 * Explanation: There is a cycle in the linked list, where tail connects to the
 * first node.
 * 
 * 
 * 
 * 
 * Example 3:
 * 
 * 
 * Input: head = [1], pos = -1
 * Output: no cycle
 * Explanation: There is no cycle in the linked list.
 * 
 * 
 * 
 * 
 * 
 * 
 * Follow up:
 * Can you solve it without using extra space?
 * 
 */
/**
 * Definition for singly-linked list.
	* type ListNode struct {
	*     Val int
	*     Next *ListNode
	* }
 */

 // 输出环的起始位置
// 输出环的起始位置
func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	for slow != nil && fast != nil && fast.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
		if slow == fast {
			break
		}
	}
	// 典型错误，链表只有一个节点，slow和fast也相同
	if slow == fast { //在此处追上，不代表此处就是环的起始位置
		// slow和fast必须从头开始走，fast不能提前走,
		// A从相遇的位置开始走，每次一步
		// B从head开始，每次一步，相遇时即为环的第一个节点
		slow = head
		for slow != fast {
			slow, fast = slow.Next, fast.Next
		}
		return slow
	}
	return nil
}
