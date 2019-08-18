/*
 * @lc app=leetcode id=978 lang=golang
 *
 * [978] Longest Turbulent Subarray
 */
// 滑动窗口处理
// 68 ms, faster than 76.19%
// [9,9]
func maxTurbulenceSize(A []int) int {
	maxLen := 1
	left := 0
	flag := 0 // 大小符号
	for right := 1; right < len(A); right++ {
		flag = compare(A[right-1], A[right])
		if flag == 0 { // 此时是right 不满足条件
			left = right
			if right-left > maxLen {
				maxLen = right - left
			}
		} else if right == len(A)-1 || flag*compare(A[right], A[right+1]) != -1 { // 此时是right+1不满足条件
			// 有相等和符号相同两种情况
			if right-left+1 > maxLen {
				maxLen = right - left + 1
			}

			left = right
		}
	}
	return maxLen
}

func compare(a, b int) int {
	if a > b {
		return 1
	} else if a < b {
		return -1
	}
	return 0
}
