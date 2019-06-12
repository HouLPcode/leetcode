/*
 * @lc app=leetcode id=172 lang=golang
 *
 * [172] Factorial Trailing Zeroes
 */
// log时间复杂度
// 每5个数有1个5，每25个数有2个5，每125个数有3个5
// 优化后的思路，
// 		每5个数+1，即有n/5个5
// 		每25个数+1，即 +n/25个5，每25个数有2个5，其中1个5在没5个数中已经累加过了

func trailingZeroes(n int) int {
	cnt := 0 
	// *****
	for n > 0 { 
		cnt += n/5
		n /= 5
	}
	return cnt
}
