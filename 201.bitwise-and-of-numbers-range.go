/*
 * @lc app=leetcode id=201 lang=golang
 *
 * [201] Bitwise AND of Numbers Range
 */
//  20 ms 60 %
// 位运算
func rangeBitwiseAnd(m int, n int) int {
	if m == 0 {
		return 0
	}
	// 统计m的位数
	bits := uint(0)
	for i := m; i > 0; i >>= 1 {
		bits++
	}

	if n >= (1 << (bits)) { // 位数不同，直接返回0
		return 0
	}

	// 去除最高位，递归调用
	m = m & (^(1 << (bits - 1))) // 清除最高位1
	n = n & (^(1 << (bits - 1)))
	return 1<<(bits-1) + rangeBitwiseAnd(m, n)
}
