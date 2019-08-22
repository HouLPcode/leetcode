import (
	"container/heap"
)

/*
 * @lc app=leetcode id=378 lang=golang
 *
 * [378] Kth Smallest Element in a Sorted Matrix
 */
//  (44 ms)   √ Your runtime beats 54.17 %
// minHeap 实现
// 先插入首列所有数据，然后弹出一个数据，将该数据所在行的下一个数据push
// TODO 还有另一种方法，二分法实现？？？？？？？？？
type Node struct {
	R, C, Val int
}

type NodeHeap []Node

func (h NodeHeap) Len() int           { return len(h) }
func (h NodeHeap) Less(i, j int) bool { return h[i].Val < h[j].Val }
func (h NodeHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *NodeHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(Node))
}

func (h *NodeHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func kthSmallest(matrix [][]int, k int) int {
	h := &NodeHeap{}
	heap.Init(h)
	n := len(matrix)
	for i := 0; i < n; i++ { // 把第一列压入堆中
		heap.Push(h, Node{i, 0, matrix[i][0]})
	}
	v := Node{}
	for k > 0 {
		v = heap.Pop(h).(Node)
		if v.C+1 < n { // 没到列尾
			heap.Push(h, Node{v.R, v.C + 1, matrix[v.R][v.C+1]})
		}
		k--
	}
	return v.Val
}
