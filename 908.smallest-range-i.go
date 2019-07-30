/*
 * @lc app=leetcode id=908 lang=golang
 *
 * [908] Smallest Range I
 */
//  12 ms, faster than 95.45%
func smallestRangeI(A []int, K int) int {
	// 遍历数组，找到最大最小值， 计算差值和k的关系
	// len(A) == 0
	min, max := A[0], A[0]
	for _, v := range A {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	if max-min < K*2 {
		return 0
	}
	return max - min - K*2
}
