/*
 * @lc app=leetcode id=23 lang=golang
 *
 * [23] Merge k Sorted Lists
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 // 方法1: 暴力合并，每轮筛选出一个最小node加入链表，O(kkn),k表示链表数，每个链表平均n个节点
//  方法2: map缓存后排序，O(knlog(kn))
 // 方法3: 分治法，两两合并 O(knlogk)
func mergeKLists(lists []*ListNode) *ListNode {
    if len(lists) == 0{
		return nil
	}else if len(lists) == 1{
		return lists[0]
	}else if len(lists) == 2{
		return mergeTwoLists(lists[0], lists[1])
	}else{
		l1 := mergeKLists(lists[:len(lists)/2])
		l2 := mergeKLists(lists[len(lists)/2:])
		return mergeTwoLists(l1, l2)
	}
	return nil // 不会到这一步
}
// copy leetcode 21
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    if l1 == nil {
		return l2
	}else if l2 == nil{
		return l1
	}
	var head, cur *ListNode
	if l1.Val < l2.Val{
		head, cur = l1, l1
		l1 = l1.Next
   	}else {
		head, cur = l2, l2
		l2 = l2.Next
   	}

	for l1!=nil && l2 != nil {
		if l1.Val < l2.Val{
			 cur.Next =  l1 
			 cur = cur.Next
			 l1 = l1.Next
		}else {
			cur.Next = l2
			cur= cur.Next
			l2 = l2.Next
		}
	}
	if l1 != nil{
		cur.Next = l1
	}else {
		cur.Next = l2
	}
	return head
}

