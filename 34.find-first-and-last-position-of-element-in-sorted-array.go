/*
 * @lc app=leetcode id=34 lang=golang
 *
 * [34] Find First and Last Position of Element in Sorted Array
 */
//8 ms 89.55 %
// 第一次二分查找左边界，第二次从该边界值开始查找
func searchRange(nums []int, target int) []int {
	rtn := []int{-1, -1}
	if len(nums) == 0 { //一定注意空数组的处理，可能会导致循环推出后nums[0]访问越界
		return rtn
	}
	s, e := 0, len(nums)-1
	for s < e { // 找左边界
		mid := (e-s)/2 + s
		if nums[mid] >= target {
			e = mid
		} else {
			s = mid + 1
		}
	}
	if nums[s] != target {
		return rtn
	}
	rtn[0] = s
	e = len(nums) - 1
	for s+1 < e { // 找右边界
		mid := (e-s)/2 + s
		if nums[mid] > target {
			e = mid - 1
		} else {
			// ==
			s = mid
		}
	}
	if nums[e] == target {
		rtn[1] = e
	} else {
		rtn[1] = s
	}
	return rtn
}
