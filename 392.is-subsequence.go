/*
 * @lc app=leetcode id=392 lang=golang
 *
 * [392] Is Subsequence
 */
// 4 ms 100 %
//  两个指针直接遍历
func isSubsequence(s string, t string) bool {
	tl := len(t)
	sl := len(s)
	if sl > tl {
		return false
	}
	// 遍历s，每个元素从t中找一个
	i, j := 0, 0
	for ; i < len(s); i++ {
		for ; j < len(t) && t[j] != s[i]; j++ {
		}
		if j == len(t) {
			return false
		}
		j++ // 一定注意，不要忘了此处的j需要后移一位
	}
	return i == len(s)
}
