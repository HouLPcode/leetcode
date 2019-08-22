import "container/heap"

/*
 * @lc app=leetcode id=373 lang=golang
 *
 * [373] Find K Pairs with Smallest Sums
 */
// 16 ms, faster than 75.56%
// maxHeap
type Num struct {
	V1, V2 int
}

type MaxHeap []Num

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i].V1+h[i].V2 > h[j].V1+h[j].V2 }
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

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	maxH := new(MaxHeap)
	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); j++ { // 所有的组合入堆
			if maxH.Len() < k { // 维持K 大堆
				heap.Push(maxH, Num{nums1[i], nums2[j]})
			} else if nums1[i]+nums2[j] < (*maxH)[0].V1+(*maxH)[0].V2 {
				heap.Pop(maxH)
				heap.Push(maxH, Num{nums1[i], nums2[j]})
			}
		}
	}
	rtn := make([][]int, min(k, len(nums1)*len(nums2)))
	for i := len(rtn) - 1; i >= 0; i-- {
		d := heap.Pop(maxH).(Num)
		rtn[i] = []int{d.V1, d.V2}
	}
	return rtn
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
