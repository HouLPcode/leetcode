/*
 * @lc app=leetcode id=69 lang=golang
 *
 * [69] Sqrt(x)
 */
//  0 ms
// 折半查找
func mySqrt(x int) int {
	if x == 0 || x == 1 { // 题目中指定x非负
		return x
	}
	left, right := 1, x
	for left <= right {
		mid := (right-left)/2 + left
		if mid > x/mid { //注意：采用除法，防止 mid*mid 越界
			right = mid - 1
		} else if mid <= x/mid {
			left = mid + 1
		}
	}
	return left - 1
}
