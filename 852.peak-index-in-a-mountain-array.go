/*
 * @lc app=leetcode id=852 lang=golang
 *
 * [852] Peak Index in a Mountain Array
 */
//  8 ms 96.77 %
func peakIndexInMountainArray(A []int) int {
	// 二分查找，题目中给出i一定存在
	mid := 0
	for l,r := 0, len(A)-1; l <= r; {
		mid = (r-l)/2 + l
		if A[mid]>A[mid-1] && A[mid]>A[mid+1] {// mid存在，所以不越界
			break
		} else if A[mid] > A[mid-1] && A[mid] < A[mid + 1] {
			l = mid + 1
		} else {
			r = mid -1
		}
	}
	return mid
}

