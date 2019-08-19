/*
 * @lc app=leetcode id=551 lang=golang
 *
 * [551] Student Attendance Record I
 */
// 0 ms
// "LALL"
// "PPALLL"
// "PPALLP"
// 注意读清题目，是连续两个L
func checkRecord(s string) bool {
	A := 0
	for i := 0; i < len(s); i++ {
		if s[i] == 'A' {
			A++
			if A == 2 {
				return false
			}
		} else if s[i] == 'L' {
			if i+2 < len(s) && s[i+1] == 'L' && s[i+2] == 'L' {
				return false
			}
		}
	}
	return true
}
