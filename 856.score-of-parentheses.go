/*
 * @lc app=leetcode id=856 lang=golang
 *
 * [856] Score of Parentheses
 */
// (0 ms)
// 左括号当0
// 有括号的时候出栈两个元素，入栈 top1*2+top2
// 初始栈加入一个元素，避免出栈越界，最终结果在stack[0]中
func scoreOfParentheses(S string) int {
	stack := []int{0} // 先房间去一个0，避免pop的时候出错
	for i := 0; i < len(S); i++ {
		if S[i] == '(' { // (=0
			stack = append(stack, 0)
		} else {
			top1 := stack[len(stack)-1]
			top2 := stack[len(stack)-2] // 不会越界
			stack = stack[:len(stack)-2]
			if top1 == 0 { // 栈顶(
				stack = append(stack, 1+top2) // 注意此处是top2 + 1
			} else { // 栈顶 num
				stack = append(stack, top1*2+top2) // 不会出现连续多个数字在栈中？？？
			}
		}
	}
	return stack[0]
}
