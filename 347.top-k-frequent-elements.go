/*
 * @lc app=leetcode id=347 lang=golang
 *
 * [347] Top K Frequent Elements
 */
type Data struct{
	Key int
	Cnt int
}

 type MaxIntHeap []Data

 func (h MaxIntHeap) Len() int           { return len(h) }
 //此处编译错误 non-integer slice index i
 func (h MaxIntHeap) Less(i, j Data) bool { return h[i].Cnt.(int) > h[j].Cnt.(int) }
 func (h MaxIntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
 
 func (h *MaxIntHeap) Push(x interface{}) {
	 *h = append(*h, x.(Data))
 }
 
 func (h *MaxIntHeap) Pop() interface{} {
	 old := *h
	 n := len(old)
	 x := old[n-1]
	 *h = old[0 : n-1]
	 return x
 }

func topKFrequent(nums []int, k int) []int {
	// 1. 遍历数组，统计频次
	numsmap := make(map[int]int)
	for _,v := range nums{
		numsmap[v]++
	}
	// 2. 构建最大堆(堆结构怎么存储)
	h := &MaxIntHeap{}
	heap.Init(h)
	for k,v := range numsmap{
		heap.Push(h, Data{Key:k,Cnt:v})
	}
	
	// 3. 每次弹出一个，弹出k次
	rnt := []int{}
	for i:=0; i<k; i++{
		rnt = append(rnt,heap.pop(h))
	}
	return rnt
}

