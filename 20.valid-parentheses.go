/*
 * @lc app=leetcode id=20 lang=golang
 *
 * [20] Valid Parentheses
 *
 * https://leetcode.com/problems/valid-parentheses/description/
 *
 * algorithms
 * Easy (36.32%)
 * Total Accepted:    569.2K
 * Total Submissions: 1.6M
 * Testcase Example:  '"()"'
 *
 * Given a string containing just the characters '(', ')', '{', '}', '[' and
 * ']', determine if the input string is valid.
 * 
 * An input string is valid if:
 * 
 * 
 * Open brackets must be closed by the same type of brackets.
 * Open brackets must be closed in the correct order.
 * 
 * 
 * Note that an empty string is also considered valid.
 * 
 * Example 1:
 * 
 * 
 * Input: "()"
 * Output: true
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: "()[]{}"
 * Output: true
 * 
 * 
 * Example 3:
 * 
 * 
 * Input: "(]"
 * Output: false
 * 
 * 
 * Example 4:
 * 
 * 
 * Input: "([)]"
 * Output: false
 * 
 * 
 * Example 5:
 * 
 * 
 * Input: "{[]}"
 * Output: true
 * 
 * 
 */
func isValid(s string) bool {
	// 通过slice实现栈操作
	stack := make([]rune, 0)//必须指明长度，0必须写
	m := map[rune]rune{ //左括号是val，右括号是key
		')': '(',
		']': '[',
		'}': '{', //分段显示，此处必须有逗号
	}

	for _, v := range s {
		if _, ok := m[v]; !ok { // 注意是;不是,
			// 左边括号
			stack = append(stack, v)
		} else if len(stack) != 0 && stack[len(stack)-1] == m[v] { //此处是m[v]，不是v，栈中只存储左括号
			// 右括号
			stack = stack[:len(stack)-1] //出栈
		} else {
			return false
		}
	}
	return len(stack) == 0
}
