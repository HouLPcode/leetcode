/*
 * @lc app=leetcode id=4 lang=golang
 *
 * [4] Median of Two Sorted Arrays
 */
 // err-----------------------------------------
//  [1,2]\n[3,4]
 // 二分查找注意事项
//  1. 循环条件有没有等号
//  2. 中位数判断，左，右，两个都有
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	l1 := len(nums1)
	l2 := len(nums2)
	if l1 > l2 {
		return findMedianSortedArrays(nums2, nums1) // 二分查找短的那个数组，起到加速的作用
	}
	k := (l1 + l2 + 1) / 2 // 左中位数，表示第几个数，不是下标
	// 二分遍历nums1
	s , e := 0, l1 - 1
	for s <= e { // 没有等于？？？
		// 暂时不考虑不取nums1和全取nums1两种情况
		mid1 := (e - s) / 2 + s // 下标
		mid2 := k - (mid1+1) - 1 // 下标？？？这样写mid2仍有可能是赋值
		if (mid1 < l1-1 && mid2 < l2-1 && nums1[mid1] < nums2[mid2+1] && nums2[mid2] < nums1[mid1+1]) {
			if (l1 + l2 ) & 1 == 0 {
				// 偶数
				return float64(max(nums1[mid1], nums2[mid2]) + min(nums1[mid1+1], nums2[mid2+1])) / 2.0
			}else {
				return float64(max(nums1[mid1], nums2[mid2]))
			}
		} else if (mid2 < l2-1 && mid1 == l1-1 && nums1[mid1] < nums2[mid2+1]) {
			if (l1 + l2 ) & 1 == 0 {
				// 偶数
				return float64(max(nums1[mid1], nums2[mid2]) + nums2[mid2+1]) / 2.0
			}else {
				return float64(max(nums1[mid1], nums2[mid2]))
			}
		} else if (mid1 < l1-1 && mid2 == l2-1&& nums2[mid2] < nums1[mid1+1]) { // +1 可能导致nums1或nums2越界
			// 找到满足要求的mid1
			if (l1 + l2 ) & 1 == 0 {
				// 偶数
				return float64(max(nums1[mid1], nums2[mid2]) + nums1[mid1+1]) / 2.0
			}else {
				return float64(max(nums1[mid1], nums2[mid2]))
			}
		} else if mid1 == l1-1 && mid2 == l2-1 {
			return float64(nums1[mid1] + nums2[mid2]) / 2.0 
		} else if nums1[mid1] > nums2[mid2+1] {
			e = mid1 - 1
		} else if nums1[mid1+1] < nums2[mid2] {
			s = mid1 + 1
		}
	}
	if e < 0 {
		// nums1没用
		if (l1 + l2 ) & 1 == 0 {
			return float64(nums2[k-1] + nums2[k]) / 2.0
		}
		return float64(nums2[k-1])
	} else if s == l1 {
		// nums1全用
		if (l1 + l2 ) & 1 == 0 {
			return float64(nums2[k-l1-1]+nums2[k-l1]) / 2.0
		}
		return float64(nums2[k-l1-1])
	}
	return 0.0
}

func max(a , b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

