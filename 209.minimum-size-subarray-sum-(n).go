/*
 * @lc app=leetcode id=209 lang=golang
 *
 * [209] Minimum Size Subarray Sum
 */
//  8 ms 36.43 % O(n^2)???????????????????
//  O(n),待优化，  O(n log n)
// 从左往右，双指针遍历
func minSubArrayLen(s int, nums []int) int {
	sum := 0
	minlen := len(nums) + 1 //
	i := 0
	j := 0
	for j < len(nums) {
		sum += nums[j]
		if sum >= s {
			for sum >= s { // 注意此处先移动i，之后才统计长度，且长度为当前窗口长度+1
				sum -= nums[i]
				i++
			}
			if minlen > j-i+1+1 {
				minlen = j - i + 1 + 1
			}
		}
		j++
	}
	if minlen == len(nums)+1 { // 注意，minlen初始化的时候不能为len(nums),否则无法区分返回0还是返回len(nums)
		return 0 // 注意，所有元素和不够s
	}
	return minlen
}
