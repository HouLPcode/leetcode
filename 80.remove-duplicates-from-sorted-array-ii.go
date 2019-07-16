/*
 * @lc app=leetcode id=80 lang=golang
 *
 * [80] Remove Duplicates from Sorted Array II
 */
// 4 ms, faster than 94.18%
func removeDuplicates(nums []int) int {
	i := 0 // 保证i前面的数字不超过2个，满足题目中要求
	for j := 0; j < len(nums); j++ {
		if i < 2 || nums[j] > nums[i-2] {
			nums[i] = nums[j]
			i++
		}
	}
	return i
}
