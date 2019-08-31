/*
 * @lc app=leetcode id=821 lang=golang
 *
 * [821] Shortest Distance to a Character
 */
// 0 ms
// 遍历两遍
func shortestToChar(S string, C byte) []int {

	rtn := make([]int, len(S))

	// 左往右
	pri := -len(S) // 前一个C的位置, 初始为-len(S), 则rtn中记录的值肯定比右往左的时候大
	for i := 0; i < len(S); i++ {
		if S[i] == C {
			pri = i
		}
		rtn[i] = i - pri
	}

	// 右往左
	next := len(S) * 2 // 2倍长度，保证右往左的时候最大
	for i := len(S) - 1; i >= 0; i-- {
		if S[i] == C {
			next = i
		}
		if rtn[i] > next-i {
			rtn[i] = next - i
		}
	}
	return rtn
}
