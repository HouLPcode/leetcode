/*
 * @lc app=leetcode id=493 lang=golang
 *
 * [493] Reverse Pairs
 */
//  超时
func reversePairs(nums []int) int {
	// 暴力搜索 O(n^2)
	rtn := 0
	l := len(nums)
	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			if nums[i] > 2*nums[j] {
				rtn++
			}
		}
	}
	return rtn
}
