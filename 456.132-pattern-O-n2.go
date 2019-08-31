import "math"

/*
 * @lc app=leetcode id=456 lang=golang
 *
 * [456] 132 Pattern
 */
// 360 ms, faster than 13.89%
// 时间O(n^2)
// [1,0,1,-4,-3]
func find132pattern(nums []int) bool {
	// 遍历一遍中间值j，同时记录一直的j最小值作为i
	num_i := math.MaxInt64
	for j := 0; j < len(nums); j++ {
		if nums[j] < num_i {
			num_i = nums[j]
		} else {
			// 从j+1开始遍历，找到 > 的值就是k
			for k := j + 1; k < len(nums); k++ {
				if nums[j] > nums[k] && nums[k] > num_i {
					return true
				}
			}
		}
	}
	return false
}
