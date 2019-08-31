/*
 * @lc app=leetcode.cn id=921 lang=golang
 *
 * [921] 使括号有效的最少添加
 */
// 0 ms
func minAddToMakeValid(S string) int {
	rtn := 0
	stack := 0 // 左括号的数量，没必要真的保存起来，计数就行
	for i := 0; i < len(S); i++ {
		if S[i] == '(' {
			stack++
		} else {
			if stack == 0 {
				rtn++
			} else {
				stack--
			}
		}
	}
	return rtn + stack
}
