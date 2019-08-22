/*
 * @lc app=leetcode id=905 lang=golang
 *
 * [905] Sort Array By Parity
 */
// (52 ms) √ Your runtime beats 91.99 %
func sortArrayByParity(A []int) []int {
	// 偶数在左，奇数在右
	r := len(A) - 1 // r右边都是奇数，且不包括r
	for i := 0; i <= r; {
		if A[i]&1 == 1 {
			A[i], A[r] = A[r], A[i]
			r--
		} else {
			i++
		}
	}
	return A
}
