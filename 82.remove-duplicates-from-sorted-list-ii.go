/*
 * @lc app=leetcode id=82 lang=golang
 *
 * [82] Remove Duplicates from Sorted List II
 */
//  12ms 8.96% 待优化
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func deleteDuplicates(head *ListNode) *ListNode {
	// 找到所有要删除节点的val值
	vals := dupVals(head)
	// 删除所有val值的节点
	for _,v := range vals{
		delNode(&head,v)
	}
	return head
}

func dupVals(head *ListNode) []int {
	rtn := []int{}
	vals := make(map[int]int, 0)
	for p1:=head; p1!=nil; p1=p1.Next {
		vals[p1.Val]++
	}
	for k,v := range vals {
		if v > 1{
			rtn = append(rtn, k) // 有序的属性被打乱了
		}
	}
	return rtn
}

// 注意此处指**ListNode，因为有可能删除的是头指针，需要进行head重新赋值，避免值传递的影响
func delNode(head **ListNode, val int) {
	tmpHead := &ListNode{
		Val:0,
		Next:*head,
	}
	for p1:=tmpHead; p1 != nil;  { // 一定注意，删除节点后p1不后移,否则会出现一个多余的要val节点未删除
		if p1.Next != nil && p1.Next.Val == val{ //删除节点
			p1.Next = p1.Next.Next
		}else { // 下一个节点不是要删除的节点，则p1后移。
			p1=p1.Next
		}
	}
	*head = tmpHead.Next
}
