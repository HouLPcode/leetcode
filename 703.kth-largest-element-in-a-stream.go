/*
 * @lc app=leetcode id=703 lang=golang
 *
 * [703] Kth Largest Element in a Stream
 */

// 第K大的数，用一个容量为K的小顶堆，堆顶元素即所求值,
// 从go doc中拷贝的代码
// An IntHeap is a min-heap of ints.
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

// This example inserts several ints into an IntHeap, checks the minimum,
// and removes them in order of priority.
// func main() {
// 	h := &IntHeap{2, 1, 5}
// 	heap.Init(h)
// 	heap.Push(h, 3)
// 	fmt.Printf("minimum: %d\n", (*h)[0])
// 	for h.Len() > 0 {
// 		fmt.Printf("%d ", heap.Pop(h))
// 	}
// }

type KthLargest struct {
	Data *MinIntHeap
	Count int
}


func Constructor(k int, nums []int) KthLargest {
	// 注意积累堆的使用方式
	h := &MinIntHeap{}
	heap.Init(h)
	for _,v := range nums{
		//堆中个数不够k个
		if h.Len() < k{
			heap.Push(h,v)
		}else{
			if v > (*h)[0]{
				heap.Pop(h)
				heap.Push(h,v)
			} 
		}
	}
	return KthLargest{
		Data:h,
		Count:k,
	}
}


func (this *KthLargest) Add(val int) int {
	if this.Data.Len() < this.Count{
		heap.Push(this.Data,val)
	}else if val > (*this.Data)[0]{//注意，data是指针，不能直接索引
		heap.Pop(this.Data)
		heap.Push(this.Data,val)
	}
	return (*this.Data)[0]
}


/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */

