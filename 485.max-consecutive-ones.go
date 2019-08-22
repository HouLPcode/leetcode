/*
 * @lc app=leetcode id=485 lang=golang
 *
 * [485] Max Consecutive Ones
 */
// 40 ms, faster than 51.22%
func findMaxConsecutiveOnes(nums []int) int {
	maxCnt := -1
	cnt := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			cnt++
		} else {
			if cnt > maxCnt {
				maxCnt = cnt
			}
			cnt = 0
		}
	}
	if cnt > maxCnt {
		maxCnt = cnt
	}
	return maxCnt
}
