/*
 * @lc app=leetcode id=153 lang=golang
 *
 * [153] Find Minimum in Rotated Sorted Array
 */
//  0 ms
//  二分法第二种情况
func findMin(nums []int) int {
	// 如果是旋转后的数组，最小值肯定小于前一个值，正常值都是大于前一个值
	s, e := 0, len(nums) // e = len
	for s < e {
		mid := (e-s)/2 + s
		if mid > 0 && nums[mid] < nums[mid-1] {
			return nums[mid]
		} else if nums[mid] > nums[0] {
			// 左有序
			s = mid + 1
		} else {
			// 右有序
			e = mid
		}
	}
	if s == len(nums) { // 没有找到想要的值，即数组没有旋转
		return nums[0]
	}
	return nums[s] // nums不为空，退出循环表示数组没有旋转
}
