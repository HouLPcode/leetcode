/*
 * @lc app=leetcode id=946 lang=golang
 *
 * [946] Validate Stack Sequences
 */
// (4 ms) √ Your runtime beats 96.23 %
// 思路一样，编码更加简练
// [1,0]\n[1,0]
func validateStackSequences(pushed []int, popped []int) bool {
	stack := []int{}
	i := 0
	for _, v := range pushed {
		stack = append(stack, v) // push
		for len(stack) > 0 && stack[len(stack)-1] == popped[i] {
			stack = stack[:len(stack)-1] // pop
			i++
		}
	}
	return i == len(popped)
}
