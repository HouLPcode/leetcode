/*
 * @lc app=leetcode id=19 lang=golang
 *
 * [19] Remove Nth Node From End of List
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// 前后两个指针，差n个节点，遍历一边，前节点即是要删除的节点
	p1,p2 := head, head
	// n从1开始
	for n > 0{ // n合理，p最多移动到nil，不回继续移动下去
		p1 = p1.Next
		n--
	}

	// 注意 删除头结点的处理
	if p1 == nil{//由于n是合理的，所以此时删除的是头结点
		head = head.Next
	}else{
		for p1.Next != nil{
			p2, p1 = p2.Next, p1.Next
		}
		p2.Next = p2.Next.Next
	}
	return head
}

