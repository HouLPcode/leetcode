/*
 * @lc app=leetcode id=349 lang=golang
 *
 * [349] Intersection of Two Arrays
 */
 func intersection(nums1 []int, nums2 []int) []int {
	// hash 处理
	// 注意去重
	nummap := make(map[int]int, 0)
	for _,v := range nums1 {
		// nummap[v]++ 通过赋值处理，而不是计数，计数可能吧一个数组中的数字返回
		nummap[v] = 1
	}
	for _,v := range nums2 {
		if nummap[v] == 1 {
			nummap[v] = 2
		}
	}

	rtn := make([]int, 0, len(nums1))
	for k,v := range nummap {
		if v == 2 {
			rtn = append(rtn, k)
		}
	}
	return rtn
}

