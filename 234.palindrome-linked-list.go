/*
 * @lc app=leetcode id=234 lang=golang
 *
 * [234] Palindrome Linked List
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
	//遍历一边得到长度，将前半部分翻转，比较翻转后的链表是否相同
	cnt := 0
	for p := head; p != nil; p = p.Next{
		cnt++
	}

	var pri,cur *ListNode = nil, head 
	for i:=0; i< cnt/2; i++{
		pri, cur, cur.Next = cur, cur.Next, pri
	}
	if cnt & 1 == 0{ // 偶数节点
		return equal(pri,cur)// pri前半部分，cur后半部分
	}
	// 奇数节点
	return equal(pri,cur.Next)// pri前半部分，cur后半部分
}
func equal(l1, l2 *ListNode)bool{
	for ;l1 != nil && l2 != nil && l1.Val == l2.Val; l1, l2 = l1.Next, l2.Next{}
	if l1 != nil || l2 != nil {
		return false
	}	
	return true
}

