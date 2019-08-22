import "container/heap"

/*
 * @lc app=leetcode id=857 lang=golang
 *
 * [857] Minimum Cost to Hire K Workers
 */
// TOP k  大堆  w/q 同质量权重越小越好
// 要求的是必须满足每个工人的最低期望工资，所以不能只看最小的几个值？？？下面的思路是错误的？？？
func mincostToHireWorkers(quality []int, wage []int, K int) float64 {
	maxH := new(MaxHeap)
	for i := 0; i < len(quality); i++ {
		if maxH.Len() < K {
			heap.Push(maxH, Person{wage[i], quality[i]})
		} else if wage[i]*(*maxH)[0].Quality < (*maxH)[0].Wage*quality[i] { // < top 此处加一次判断，否则会超时
			heap.Push(maxH, Person{wage[i], quality[i]})
			heap.Pop(maxH)
		}
	}

	// (*maxH)[0].wage / (*maxH)[0].q * sum（）
	sumQ := 0
	for i := 0; i < maxH.Len(); i++ {
		sumQ += (*maxH)[i].Quality
	}
	return float64((*maxH)[0].Wage / (*maxH)[0].Quality * sumQ)
}

type Person struct {
	Wage, Quality int
}
type MaxHeap []Person

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i].Wage*h[j].Quality > h[i].Quality*h[j].Wage }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(Person))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
