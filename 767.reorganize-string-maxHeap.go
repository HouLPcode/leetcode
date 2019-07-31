import "container/heap"

/*
 * @lc app=leetcode id=767 lang=golang
 *
 * [767] Reorganize String
 */
// (0 ms) √ Your runtime beats 100 %
// 鸽巢问题
// 大堆实现， 按照字母频次构建大堆，每次取频次最高的两个字母输出，频次-1后循环，直到堆为空
func reorganizeString(S string) string {
	// 统计S中各个字符出现的次数, 只有小写字母
	buf := make([]int, 26) // buf[字母][频次]
	for i := 0; i < len(S); i++ {
		buf[S[i]-'a']++
	}

	maxHeap := &CharsHeap{}
	heap.Init(maxHeap)
	for i := byte(0); i < 26; i++ {
		if buf[i] > 0 {
			heap.Push(maxHeap, chars{buf[i], 'a' + i})
		}
	}

	rtn := make([]byte, 0, len(S))
	for maxHeap.Len() >= 2 {
		node1 := heap.Pop(maxHeap).(chars)
		node2 := heap.Pop(maxHeap).(chars)
		if node1.cnt > (len(S)+1)/2 {
			return "" // 超过一半，不存在
		}
		rtn = append(rtn, node1.alh) // 填充顺序？？？ 会出现abba的情况吗？？？
		rtn = append(rtn, node2.alh)
		node1.cnt--
		node2.cnt--
		if node1.cnt > 0 {
			heap.Push(maxHeap, node1)
		}
		if node2.cnt > 0 {
			heap.Push(maxHeap, node2)
		}
	}
	if maxHeap.Len() == 1 {
		node := heap.Pop(maxHeap).(chars)
		if node.cnt > 1 {
			return ""
		}
		rtn = append(rtn, node.alh)
	}
	return string(rtn)
}

type chars struct {
	cnt int
	alh byte
}

// An IntHeap is a min-heap of ints.
type CharsHeap []chars

func (h CharsHeap) Len() int           { return len(h) }
func (h CharsHeap) Less(i, j int) bool { return h[i].cnt > h[j].cnt }
func (h CharsHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *CharsHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(chars))
}

func (h *CharsHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
