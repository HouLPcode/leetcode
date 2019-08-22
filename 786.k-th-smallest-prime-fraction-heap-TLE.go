import "container/heap"

/*
 * @lc app=leetcode id=786 lang=golang
 *
 * [786] K-th Smallest Prime Fraction
 */
// 大堆实现 × Time Limit Exceeded
func kthSmallestPrimeFraction(A []int, K int) []int {
	maxH := new(MaxHeap)
	for i := 0; i < len(A); i++ {
		for j := i + 1; j < len(A); j++ {
			if maxH.Len() < K {
				heap.Push(maxH, Num{A[i], A[j]})
			} else { // < top 减少一次判断
				heap.Push(maxH, Num{A[i], A[j]})
				heap.Pop(maxH)
			}
		}
	}
	return []int{(*maxH)[0].V1, (*maxH)[0].V2}
}

type Num struct {
	V1, V2 int
}
type MaxHeap []Num

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i].V1*h[j].V2 > h[i].V2*h[j].V1 }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(Num))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
