import (
	"container/heap"
	"sort"
)

/*
 * @lc app=leetcode id=692 lang=golang
 *
 * [692] Top K Frequent Words
 */

//  4 ms, faster than 100.00%

// ["the", "day", "is", "sunny", "the", "the", "the", "sunny", "is", "is"]\n4
// ["i", "love", "leetcode", "i", "love", "coding"]\n2
//  注意小堆的输出顺序
// minHeap
type Word struct {
	Str string
	Cnt int
}
type MinHeap []Word

func (h MinHeap) Len() int { return len(h) }
func (h MinHeap) Less(i, j int) bool {
	return h[i].Cnt < h[j].Cnt || (h[i].Cnt == h[j].Cnt && h[i].Str > h[j].Str) // 注意，string可以直接比较，不用strings.Compare
}
func (h MinHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(Word))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func topKFrequent(words []string, k int) []string {
	t := make(map[string]int, 0)
	for i := 0; i < len(words); i++ {
		t[words[i]]++
	}
	h := &MinHeap{}
	for key, val := range t {
		if h.Len() < k {
			heap.Push(h, Word{key, val})
		} else {
			if val > (*h)[0].Cnt || (val == (*h)[0].Cnt && key < (*h)[0].Str) { // 注意，此处一定是 *h， h会报类型错误
				heap.Pop(h)
				heap.Push(h, Word{key, val})
			}
		}
	}

	rtn := make([]Word, k)
	for i := k - 1; i >= 0; i-- { // 注意，此处逆序， 小堆， 结果需要 大到小
		rtn[i] = heap.Pop(h).(Word)
	}
	// 此处处理同频次按字母排序
	sort.Slice(rtn, func(i, j int) bool {
		return rtn[i].Cnt > rtn[j].Cnt || (rtn[i].Cnt == rtn[j].Cnt && rtn[i].Str < rtn[j].Str)
	})

	rtn1 := make([]string, k)
	for i := 0; i < k; i++ {
		rtn1[i] = rtn[i].Str
	}
	return rtn1
}
