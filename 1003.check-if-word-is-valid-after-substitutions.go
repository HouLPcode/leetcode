/*
 * @lc app=leetcode id=1003 lang=golang
 *
 * [1003] Check If Word Is Valid After Substitutions
 */
//  (0 ms)
// 通过栈进行优化
func isValid(S string) bool {
	stack := make([]byte, 0, len(S)/3) // cap瞎写的
	for i := 0; i < len(S); i++ {
		if S[i] == 'c' {
			if len(stack) > 1 && stack[len(stack)-1] == 'b' && stack[len(stack)-2] == 'a' {
				stack = stack[:len(stack)-2]
			} else {
				stack = append(stack, S[i])
			}
		} else {
			stack = append(stack, S[i])
		}
	}
	return len(stack) == 0
}
