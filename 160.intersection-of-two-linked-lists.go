/*
 * @lc app=leetcode id=160 lang=golang
 *
 * [160] Intersection of Two Linked Lists
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
    //统计A，B的长度，计算差值，然后从差值出逐步寻找直到相同的节点出现
	cntA, cntB := 0,0
	pA, pB := headA, headB
	for pA != nil{
		cntA++
		pA = pA.Next
	}
	for pB != nil{
		cntB++
		pB = pB.Next
	}
	// if cntA == 0 || cntB == 0{ // 不必须
	// 	return nil
	// }
	// 追平长度
	if cntA - cntB > 0{
		for i:=0; i< cntA - cntB; i++{
			headA = headA.Next
		}
	}else{
		for i:=0; i< cntB-cntA; i++{
			headB = headB.Next
		}
	}
	for headA != nil && headB != nil && headA != headB {
		headA = headA.Next
		headB = headB.Next
	}
	if headA == nil || headB == nil{
		return nil
	}
	return headA
}

