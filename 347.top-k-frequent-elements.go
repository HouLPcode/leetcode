/*
 * @lc app=leetcode id=347 lang=golang
 *
 * [347] Top K Frequent Elements
 */
type Data struct{
	Key int
	Cnt int
}
type ByCnt []Data

func (h ByCnt) Len() int           { return len(h) }
func (h ByCnt) Less(i, j int) bool { return h[i].Cnt > h[j].Cnt }
func (h ByCnt) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func topKFrequent(nums []int, k int) []int {
	// 1. 遍历数组，统计频次
	numsmap := make(map[int]int)
	for _,v := range nums{
		numsmap[v]++
	}

	rnt := ByCnt{}
	for k,v := range numsmap{
		rnt = append(rnt,Data{Key:k,Cnt:v})
	}

	// 2. 按照频次排序,对map排序？？？
	sort.Sort(rnt)
	// 3. 输出前k项,格式转换
	rntInt := []int{}
	for i:=0; i<k; i++{
		rntInt = append(rntInt,rnt[i].Key)
	}
	return rntInt
}

