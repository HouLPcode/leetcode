/*
 * @lc app=leetcode id=209 lang=golang
 *
 * [209] Minimum Size Subarray Sum
 */
// 典型错误，贪心思想：从首尾往中间夹，删除左右中最小的数，以保证元素个数尽可能少O(n)
func minSubArrayLen(s int, nums []int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	if sum < s {
		return 0
	}
	l, r := 0, len(nums)-1
	// for l < len(nums) && r >=0 &&
	for sum > s {
		// 一定注意，左右相同的时候，不能固定是删除左还是右
		if nums[l] > nums[r] {
			sum -= nums[r]
			r--
		} else if nums[l] < nums[r] {
			sum -= nums[l]
			l++
		}
	}
	return r - l + 1 + 1
}
