import (
    "container/heap"
    "fmt"
)

/*
 * @lc app=leetcode id=973 lang=golang
 *
 * [973] K Closest Points to Origin
 */
// 小堆尝试， 代码未写完
func kClosest(points [][]int, K int) [][]int {
    // K大堆实现
    ph := &PointHeap{}
    heap.Init(ph)
    for i:=0; i<len(points); i++ {
        if ph.Len() < K {
            heap.Push(ph,points[i])
        } else {
            // 堆顶元素怎么获取？？？
            if ph[K-1] > points[i][0]*points[i][0]+points[i][1]*points[i][1]{
                heap.Pop(ph)   
                heap.Push(ph,points[i])
            }
        }
    }
    return [][]int(*ph)
}

// An IntHeap is a min-heap of ints.
type PointHeap [][]int

func (h PointHeap) Len() int           { return len(h) }
func (h PointHeap) Less(i, j int) bool { return h[i][0]*h[i][0] + h[i][1]*h[i][1] < h[j][0]*h[j][0] + h[j][1]*h[j][1] }
func (h PointHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *PointHeap) Push(x interface{}) {
    // Push and Pop use pointer receivers because they modify the slice's length,
    // not just its contents.
    *h = append(*h, x.([]int))
}

func (h *PointHeap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[0 : n-1]
    return x
}
