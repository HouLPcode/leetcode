/*
 * @lc app=leetcode id=445 lang=golang
 *
 * [445] Add Two Numbers II
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 12 ms, faster than 73.33%
// 法1 栈空间
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	stack1 := []int{}
	for p := l1; p != nil; p = p.Next {
		stack1 = append(stack1, p.Val)
	}
	stack2 := []int{}
	for p := l2; p != nil; p = p.Next {
		stack2 = append(stack2, p.Val)
	}
	i, j := len(stack1)-1, len(stack2)-1
	dummy := ListNode{
		0,
		nil,
	}
	flag := 0 //进位标值
	for i >= 0 && j >= 0 {
		// 头插法
		newNode := ListNode{}
		newNode.Val = (flag + stack1[i] + stack2[j]) % 10
		newNode.Next = dummy.Next
		dummy.Next = &newNode
		flag = (flag + stack1[i] + stack2[j]) / 10
		i--
		j--
	}
	for i >= 0 {
		newNode := ListNode{}
		newNode.Val = (flag + stack1[i]) % 10
		newNode.Next = dummy.Next
		dummy.Next = &newNode
		flag = (flag + stack1[i]) / 10
		i--
	}
	for j >= 0 {
		newNode := ListNode{}
		newNode.Val = (flag + stack2[j]) % 10
		newNode.Next = dummy.Next
		dummy.Next = &newNode
		flag = (flag + stack2[j]) / 10
		j--
	}
	if flag != 0 {
		newNode := ListNode{}
		newNode.Val = 1
		newNode.Next = dummy.Next
		dummy.Next = &newNode
	}
	return dummy.Next
}
