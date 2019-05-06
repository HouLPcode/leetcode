/*
 * @lc app=leetcode id=69 lang=golang
 *
 * [69] Sqrt(x)
 */
func mySqrt(x int) int {
	// 开方的解题思路
	// 方法1： y = x^2 是单调函数，可以从 1->n 二分遍历
	// 题目中指定x非负
	if x == 0 || x == 1{
		return x
	}
	left,right := 1,x
	for left <= right{
		mid := (left + right) / 2
		if mid == x / mid { //注意：采用除法，防止 mid*mid 越界
			return mid
		} else if mid > x / mid{
			right = mid -1
		}else if mid < x / mid{
			left = mid + 1
		}
	}
	// 跳出循环时，left^2 > x, right^2 < x,取较小值
	return right
}

