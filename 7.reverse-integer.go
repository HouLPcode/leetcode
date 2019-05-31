/*
 * @lc app=leetcode id=7 lang=golang
 *
 * [7] Reverse Integer
 */
import "math"

// 注意，负数不用单独处理
func reverse(x int) int {
	rev := 0
	pop := 0
	for x != 0 {
		pop = x % 10 // -3%10 = -3
		x /= 10       // -123 / 10 = -12
		if rev > math.MaxInt32/10 || (rev == math.MaxInt32/10 && pop > 7) {
			return 0
		}
		if rev < math.MinInt32/10 || (rev == math.MinInt32/10 && pop < -8) {
			return 0
		}
		rev = rev*10 + pop
	}
	return rev
}
