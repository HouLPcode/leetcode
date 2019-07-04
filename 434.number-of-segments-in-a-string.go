/*
 * @lc app=leetcode id=434 lang=golang
 *
 * [434] Number of Segments in a String
 */
//  0 ms
// 统计首字母个数，而不是空格个数
// 统计空格个数比较麻烦，需要清楚首尾空格，单词书是空格数+1,且末尾空格不能计算
func countSegments(s string) int {
	cnt := 0
	for i := 0; i < len(s); {
		if s[i] != ' ' {
			cnt++
			for i < len(s) && s[i] != ' ' {
				i++
			}
		} else {
			i++
		}
	}
	return cnt
}

