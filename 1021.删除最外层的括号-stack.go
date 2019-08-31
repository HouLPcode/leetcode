/*
 * @lc app=leetcode.cn id=1021 lang=golang
 *
 * [1021] 删除最外层的括号
 */
// 0 ms
func removeOuterParentheses(S string) string {
	stack := [][]byte{}
	tmp := []byte{}
	l := 0 // 一组的起始位置
	for i := 0; i < len(S); i++ {
		if S[i] == '(' {
			tmp = append(tmp, S[i])
		} else {
			tmp = tmp[:len(tmp)-1]
			if len(tmp) == 0 {
				stack = append(stack, []byte(S[l+1:i]))
				l = i + 1
			}
		}
	}
	rtn := make([]byte, 0, len(S))
	for i := 0; i < len(stack); i++ {
		rtn = append(rtn, stack[i]...)
	}
	return string(rtn)
}
