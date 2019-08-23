/*
 * @lc app=leetcode id=674 lang=golang
 *
 * [674] Longest Continuous Increasing Subsequence
 */
//(8 ms) √ Your runtime beats 93 %
func findLengthOfLCIS(nums []int) int {
	// 两个指针组成窗口进行判断
	if len(nums) == 0 {
		return 0
	}
	// else if len(nums) == 1 {
	// 	return 1
	// }
	s, e := 0, 1 // s和e指向不同元素
	lens := 0
	for e < len(nums) {
		if nums[e-1] < nums[e] {
			e++
		} else {
			if e-s > lens {
				lens = e - s
			}
			s = e
			e++
		}
	}
	if e-s > lens {
		lens = e - s
	}
	return lens
}
