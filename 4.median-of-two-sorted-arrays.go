/*
 * @lc app=leetcode id=4 lang=golang
 *
 * [4] Median of Two Sorted Arrays
 */
// (12 ms) √ Your runtime beats 84.47 %
//  [1,2]\n[3,4]
// 二分查找注意事项
//  1. 循环条件有没有等号
//  2. 中位数判断，左，右，两个都有
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	len1 := len(nums1)
	len2 := len(nums2)
	if len1 > len2 { // to ensure len1<=len2  // 这样保证时间复杂度 O(log min(len1,len2))
		nums1, nums2 = nums2, nums1 // 二分查找短的那个数组，起到加速的作用
		len1, len2 = len2, len1
	}
	left, right := 0, len1
	halfLen := (len1 + len2 + 1) / 2 //节点的个数，奇数时左节点在左半部分
	for left <= right {
		mid1 := (right-left)/2 + left                    // 右半边下标
		mid2 := halfLen - mid1                           // 右半边下标
		if mid1 < right && nums1[mid1] < nums2[mid2-1] { // mid1太小
			left = mid1 + 1 // mid1 is too small
		} else if mid1 > left && nums1[mid1-1] > nums2[mid2] { // mid1太大
			right = mid1 - 1 // mid1 is too big
		} else { // mid1 is perfect  找到正确的mid1
			maxLeft := 0
			if mid1 == 0 { // nums1数组全在右边
				maxLeft = nums2[mid2-1]
			} else if mid2 == 0 { // nums1 数组全在左边
				maxLeft = nums1[mid1-1]
			} else { // 左右都有，中间的最大值
				maxLeft = max(nums1[mid1-1], nums2[mid2-1])
			}
			if (len1+len2)&1 == 1 { // 总个数为奇数，返回中位数
				return float64(maxLeft)
			}

			minRight := 0
			if mid1 == len1 { // nums1 全在左半边
				minRight = nums2[mid2]
			} else if mid2 == len2 { // nums2全在左半边
				minRight = nums1[mid1]
			} else { // 左右都有，中间值最小的
				minRight = min(nums2[mid2], nums1[mid1])
			}
			// 总数为偶数，返回中间两数平均值
			return float64(maxLeft+minRight) / 2.0
		}
	}
	return 0.0
}

func max(a, b int) int {
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
