/*
 * @lc app=leetcode id=264 lang=golang
 *
 * [264] Ugly Number II
 */
//  0 ms
// 法1： int32有1690个丑数，可以提前计算出来，然后O(1)查表直接返回
// 法2： 循环计算 以下实现
func nthUglyNumber(n int) int {
	// n >= 1
	buf := []int{1}
	index2, index3, index5 := 0, 0, 0
	for i := 0; i < n; i++ {
		next := min(buf[index2]*2, buf[index3]*3, buf[index5]*5)
		buf = append(buf, next)
		if next == buf[index2]*2 {
			index2++
		}
		if next == buf[index3]*3 {
			index3++
		}
		if next == buf[index5]*5 {
			index5++
		}
	}
	return buf[n-1]
}

func min(a, b, c int) int {
	if a <= b && a <= c {
		return a
	} else if b <= a && b <= c {
		return b
	}
	return c
}
