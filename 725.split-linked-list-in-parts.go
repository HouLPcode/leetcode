/*
 * @lc app=leetcode id=725 lang=golang
 *
 * [725] Split Linked List in Parts
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 0 ms 时间O(n), 空间O(k)，存储了k个节点的个数
func splitListToParts(root *ListNode, k int) []*ListNode {
	// 先获取链表节点的个数
	cnt := 0
	for p := root; p != nil; p = p.Next {
		cnt++
	}
	rtn := make([]*ListNode, 0, k)
	if cnt < k { // 节点不够
		for p := root; p != nil; {
			rtn = append(rtn, p)
			p, p.Next = p.Next, nil
		}
		for i := len(rtn); i < k; i++ {
			rtn = append(rtn, nil)
		}
	} else {
		// 节点够
		nodes := make([]int, k)
		node := cnt / k
		for i := 0; i < k; i++ {
			nodes[i] = node
		}
		for i := 0; i < cnt%k; i++ { // 不够整除的节点均摊到前面节点中
			nodes[i]++
		}

		p := root
		for _, v := range nodes {
			rtn = append(rtn, p)
			for i := 0; i < v-1; i++ { // 少遍历一次，用来将链表断开
				p = p.Next
			}
			p, p.Next = p.Next, nil
		}
	}
	return rtn
}
