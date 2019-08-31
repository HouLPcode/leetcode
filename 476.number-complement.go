/*
 * @lc app=leetcode id=476 lang=golang
 *
 * [476] Number Complement
 */
// 0 ms
func findComplement(num int) int {
	bits := uint(0)
	for num0 := num; num0 != 0; num0 >>= 1 {
		bits++
	}
	return num ^ (1<<bits - 1)
}
