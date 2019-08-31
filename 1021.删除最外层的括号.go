/*
 * @lc app=leetcode.cn id=1021 lang=golang
 *
 * [1021] 删除最外层的括号
 */
// 0 ms
// 不用栈，直接计数处理
func removeOuterParentheses(S string) string {
	rtn := make([]byte, 0, len(S))
	opened := 0 // 统计左括号的数量
	for i := 0; i < len(S); i++ {
		if S[i] == '(' {
			if opened > 0 { // i是第一个左括号，直接放弃添加
				rtn = append(rtn, S[i])
			}
			opened++ // 更新左括号数量
		} else {
			if opened > 1 { // 不是最后一个右括号
				rtn = append(rtn, S[i])
			}
			opened--
		}
	}
	return string(rtn)
}
