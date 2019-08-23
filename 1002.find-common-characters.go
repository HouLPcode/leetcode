/*
 * @lc app=leetcode id=1002 lang=golang
 *
 * [1002] Find Common Characters
 */
// (0 ms)
func commonChars(A []string) []string {
	// 统计字母出现次数
	chars := make([]int, 26)
	for i := 0; i < len(A[0]); i++ { // 统计第一个字符串中每个字符的顺序
		chars[A[0][i]-'a']++
	}
	// 则最终结果只能是chars的子集
	tmp := make([]int, 26)
	for i := 1; i < len(A); i++ {
		for j := 0; j < len(A[i]); j++ { // 统计第一个字符串中每个字符的顺序
			tmp[A[i][j]-'a']++
		}
		for i := 0; i < 26; i++ { // 取交集
			if tmp[i] < chars[i] {
				chars[i] = tmp[i]
			}
			tmp[i] = 0 // 清0， 供下一个字符串统计使用
		}
	}
	rtn := []string{}
	for k, v := range chars {
		for v > 0 {
			rtn = append(rtn, string(k+'a'))
			v--
		}
	}
	return rtn
}
