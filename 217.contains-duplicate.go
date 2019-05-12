/*
 * @lc app=leetcode id=217 lang=golang
 *
 * [217] Contains Duplicate
 */
func containsDuplicate(nums []int) bool {
	//法1 先排序，后查重
	// 法2 map
	buf := make(map[int]bool)
	for _,v := range nums{
		if _,ok := buf[v]; ok{
			return true
		}
		buf[v] = true
	}
	return false
}

