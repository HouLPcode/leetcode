/*
 * @lc app=leetcode id=15 lang=golang
 *
 * [15] 3Sum
 */
func threeSum(nums []int) [][]int {
	// map 缓存
	// 输入数组中可以有重复元素
	// 输出结果不能冗余
	rtn := [][]int{}
	bufmap := make(map[int]bool)
	for _,v := range nums{
		bufmap[v] = true
	}

	for k1,v1 := range nums{
		for k2,v2 := range nums{
			if k1 == k2{
				continue
			}
			if _,ok := bufmap[-v1-v2]; ok{
				rtn = append(rtn,[]int{v1,v2,-v1-v2})
			}
		}
	}
	return rtn
}

