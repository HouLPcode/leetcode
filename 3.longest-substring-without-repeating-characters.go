/*
 * @lc app=leetcode id=3 lang=golang
 *
 * [3] Longest Substring Without Repeating Characters
 */
// (0 ms) √ Your runtime beats 100 %
// 两个指针组成子序列，map中缓存字母出现的位置，不是出现过没
// ✘ testcase: '"abba"'
// " " // 注意，不是只有小写字母的
func lengthOfLongestSubstring(s string) int {
	buf := make([]int, 128) // 存储每个字符出现的位置
	for i := 0; i < 128; i++ {
		buf[i] = -1
	}
	maxLen := 0
	left, right := 0, 0
	for ; right < len(s); right++ {
		if buf[int(s[right])] == -1 {
			buf[int(s[right])] = right // 没有出现过
		} else { // 出现过
			if right-left > maxLen {
				maxLen = right - left
			}
			// 注意：更新l坐标的时候，需要把l移动过程中出现的元素全部还原
			i := left
			for ; i < buf[int(s[right])]+1; i++ {
				buf[int(s[i])] = -1
			}
			left = i
			buf[int(s[right])] = right // 注意，不要忘了在此处更新该字符出现的位置
		}
	}
	if right-left > maxLen {
		maxLen = right - left
	}

	return maxLen
}
