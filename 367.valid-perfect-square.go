/*
 * @lc app=leetcode id=367 lang=golang
 *
 * [367] Valid Perfect Square
 */
func isPerfectSquare(num int) bool {
	// 方法1: 二分法
	if num == 0 || num == 1{
		return true
	}
	left, right := 1, num
	for left <= right {
		mid := (left + right) / 2
		if mid == num / mid && num % mid == 0{
			return true
		}else if mid > num / mid{
			right = mid -1
		}else {
			left = mid + 1
		}
	}
	return false
}

