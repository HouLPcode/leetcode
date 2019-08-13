/*
 * @lc app=leetcode id=5 lang=golang
 *
 * [5] Longest Palindromic Substring
 */
//  (0 ms)
// 法1. 中心扩散
// 法2. dp
// 以下实现中心扩散
func longestPalindrome(s string) string {
	if s == "" {
		return ""
	}
	maxLen := 0
	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		l1 := expandAroundCenter(s, i, i)
		l2 := expandAroundCenter(s, i, i+1)
		if l1 > maxLen {
			maxLen = l1
			start = i - l1/2 // 起止坐标处理
			end = i + l1/2
		}
		if l2 > maxLen {
			maxLen = l2
			start = i - l2/2 + 1
			end = i + l2/2
		}
	}
	return s[start : end+1]
}

// 返回以s[l],s[r]为中心的最长回文长度
func expandAroundCenter(s string, l, r int) int {
	for l >= 0 && r < len(s) && s[l] == s[r] {
		// 左右扩散
		l--
		r++
	}
	return r - l - 1 // 注意退出的时候，l， r 多扩散了一次
}
