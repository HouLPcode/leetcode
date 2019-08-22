/*
 * @lc app=leetcode id=896 lang=golang
 *
 * [896] Monotonic Array
 */
// (64 ms) √ Your runtime beats 75.89 %
// 通过两个函数，分别判断是否单调，不要想的太复杂了
// 另一种思路：遍历一遍，记录中间的增减标记
func isMonotonic(A []int) bool {
	return isIncrease(A) || isDec(A)
}

func isIncrease(A []int) bool { // 递增
	for i := 0; i+1 < len(A); i++ {
		if A[i] <= A[i+1] {

		} else {
			return false
		}
	}
	return true
}

func isDec(A []int) bool { // 递减
	for i := 0; i+1 < len(A); i++ {
		if A[i] >= A[i+1] {

		} else {
			return false
		}
	}
	return true
}
