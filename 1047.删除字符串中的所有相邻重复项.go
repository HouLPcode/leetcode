/*
 * @lc app=leetcode.cn id=1047 lang=golang
 *
 * [1047] 删除字符串中的所有相邻重复项
 */
//  4 ms, faster than 90.64%
// 通过栈实现
// 注意，"aaa" ==> "a", 三个重复元素剩余一个元素
func removeDuplicates(S string) string {
	stack := make([]byte, 0, len(S))
	for i := 0; i < len(S); i++ {
		if len(stack) == 0 {
			stack = append(stack, S[i])
		} else if stack[len(stack)-1] == S[i] {
			// 发现相同元素
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, S[i])
		}
	}
	return string(stack)
}
