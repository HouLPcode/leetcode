/*
 * @lc app=leetcode id=150 lang=golang
 *
 * [150] Evaluate Reverse Polish Notation
 */
 import "strconv"

func evalRPN(tokens []string) int {
	stack := make([]int, 0, len(tokens))
	for _,s := range tokens {
		switch s{
		case "+":
			stack[len(stack)-2] += stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		case "-":
			stack[len(stack)-2] -= stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		case "*":
			stack[len(stack)-2] *= stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		case "/":
			stack[len(stack)-2] /= stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		default:
			num,_ := strconv.Atoi(s)
			stack = append(stack, num)
		}
	}
	return stack[0]
}
