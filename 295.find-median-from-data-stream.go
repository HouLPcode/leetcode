import "container/heap"

// ["MedianFinder","addNum","findMedian","addNum","findMedian","addNum","findMedian","addNum","findMedian","addNum","findMedian"]\n[[],[-1],[],[-2],[],[-3],[],[-4],[],[-5],[]]
// (124 ms) √ Your runtime beats 61.11 %
// 一定注意添加元素的规则，  错误思想：先插入右堆，然后长度不等时候，右pop左push
type MedianFinder struct {
	Max *MaxIntHeap
	Min *MinIntHeap
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	return MedianFinder{new(MaxIntHeap), new(MinIntHeap)}
}

func (this *MedianFinder) AddNum(num int) {
	heap.Push(this.Max, num) //左

	if this.Max.Len() == this.Min.Len()+2 { // 高度对多差1
		heap.Push(this.Min, heap.Pop(this.Max).(int)) //右pop， 左push
	}
	if this.Min.Len() > 0 && (*(*this).Max)[0] > (*(*this).Min)[0] { // 左>右，交换一个数据，且堆大小不能变
		heap.Push(this.Min, heap.Pop(this.Max).(int)) // 左->右
		heap.Push(this.Max, heap.Pop(this.Min).(int)) // 右->左
	}
}

func (this *MedianFinder) FindMedian() float64 { // 只有取值， 不pop
	if this.Max.Len() == this.Min.Len() {
		return float64((*(*this).Max)[0]+(*(*this).Min)[0]) / 2.0 // 注意此处的写法
	}
	return float64((*(*this).Max)[0]) // odd
}

type MinIntHeap []int

func (h MinIntHeap) Len() int           { return len(h) }
func (h MinIntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinIntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinIntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *MinIntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type MaxIntHeap []int

func (h MaxIntHeap) Len() int           { return len(h) }
func (h MaxIntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxIntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxIntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *MaxIntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
