/*
 * @lc app=leetcode id=779 lang=golang
 *
 * [779] K-th Symbol in Grammar
 */
//  0 ms
// recursion 递归实现
func kthGrammar(N int, K int) int {
	if N == 1 {
		return 0
	}
	d := kthGrammar(N-1, (K+1)/2)
	if d == 0 { // 0 1
		if K&1 == 0 { // 注意K从1开始
			return 1
		} else {
			return 0
		}
	} else { // 1 0
		if K&1 == 0 {
			return 0
		} else {
			return 1
		}
	}
}
