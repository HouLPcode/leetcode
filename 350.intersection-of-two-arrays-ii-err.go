/*
 * @lc app=leetcode id=350 lang=golang
 *
 * [350] Intersection of Two Arrays II
 */
//  典型错误，不能通过长短决定那个缓存那个输出
//  [4,9,5]\n[9,4,9,8,4]
// [3,1,2]\n[1,1]
func intersect(nums1 []int, nums2 []int) []int {
	nummap := make(map[int]bool, 0)
	var rtn []int
	if len(nums1) > len(nums2) { //一定注意此处长度比较
		for _,v := range nums1 { // 缓存长的，通过短的输出结果
			nummap[v] = true
		}
		rtn = make([]int, 0, len(nums1))
		for _,v := range nums2 {
			if nummap[v] == true {
				rtn = append(rtn, v)
			}
		}
	} else {
		for _,v := range nums2 {
			nummap[v] = true
		}
		rtn = make([]int, 0, len(nums2))
		for _,v := range nums1 {
			if nummap[v] == true {
				rtn = append(rtn, v)
			}
		}
	}
	
	return rtn
}

