/*
 * @lc app=leetcode id=34 lang=golang
 *
 * [34] Find First and Last Position of Element in Sorted Array
 */
 // 方法1 二分找到一个target，然后往左右扩展，找到边界
 // 方法2 分两次分别二分找左右边界，下面是该方法的实现
 func searchRange(nums []int, target int) []int {
	rnt := []int{-1, -1}
	start, end := 0, len(nums)-1
	mid := 0
	// 分两次，分别找左边缘或右边缘
	// 同时找连个边缘不知道什么时候break？？？
	for start <= end {
		mid = (end-start)/2 + start
		if nums[mid] == target {
			if (mid == 0) || (mid > 0 && nums[mid-1] < target) {
				rnt[0] = mid // 找到左边缘
				break
			}
			end = mid - 1 //此处的mid肯定不是0？？？
		} else if nums[mid] < target {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	if rnt[0] == -1 { //没找到左边缘
		return rnt
	}
	// start = mid 有可能只有一个target值，此处不是start=mid+1
	start, end = mid, len(nums)-1 // 重新赋值，避免上轮循环的影响
	for start <= end {
		mid = (end-start)/2 + start
		if nums[mid] == target {
			if (mid == len(nums)-1) || mid < len(nums)-1 && nums[mid+1] > target {
				rnt[1] = mid //找到右边缘
				break
			}
			start = mid + 1
		} else if nums[mid] < target {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return rnt
}
