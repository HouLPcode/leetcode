/*
 * @lc app=leetcode id=169 lang=golang
 *
 * [169] Majority Element
 */
func majorityElement(nums []int) int {
	//map计数
	cnt := make(map[int]int)
	for _,v :=range nums{
		cnt[v]++
	}
	for k,v := range cnt{
		if v > len(nums)>>1{
			return k
		}
	}
	return 0//有效区域
}

