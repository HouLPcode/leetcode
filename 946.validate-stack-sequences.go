/*
 * @lc app=leetcode id=946 lang=golang
 *
 * [946] Validate Stack Sequences
 */
// 0ms
// [1,0]\n[1,0]
func validateStackSequences(pushed []int, popped []int) bool {
	if len(pushed) != len(popped) {
		return false
	} else if len(pushed) == 0 && len(popped) == 0 {
		return true
	}

	stack := []int{pushed[0]}
	i, j := 1, 0
	for {
		if len(stack) == 0 || popped[j] != stack[len(stack)-1] { // 注意此处stack不要访问-1
			if i == len(pushed) {
				return false
			}
			stack = append(stack, pushed[i])
			i++
		} else {
			stack = stack[:len(stack)-1]
			j++
			if j == len(popped) && len(stack) == 0 {
				break
			}
		}
	}
	return true
}
