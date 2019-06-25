/*
 * @lc app=leetcode id=349 lang=golang
 *
 * [349] Intersection of Two Arrays
 */
//  4 ms 94.88 %
func intersection(nums1 []int, nums2 []int) []int {
	// hash 处理
	// 注意去重
	nummap := make(map[int]bool, 0)
	for _,v := range nums1 {
		// nummap[v]++ 通过赋值处理，而不是计数，计数可能吧一个数组中的数字返回
		nummap[v] = true
	}
	for _,v := range nums2 {
		nummap[v] = true && nummap[v]
	}

	rtn := make([]int, 0, len(nums1))
	for k,v := range nummap {
		if v == true {
			rtn = append(rtn, k)
		}
	}
	return rtn
}

