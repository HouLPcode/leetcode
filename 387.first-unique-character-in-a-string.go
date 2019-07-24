/*
 * @lc app=leetcode id=387 lang=golang
 *
 * [387] First Unique Character in a String
 */
// 4 ms, faster than 100.00%
// silce替换map 加速
func firstUniqChar(s string) int {
	// 第一遍缓存每个字符出现的次数
	// 第二遍找到频次为1的索引并返回
	// 时间O(n)
	dic := [26]int{}
	for _, c := range s {
		dic[c-'a']++
	}
	for i, c := range s {
		if dic[c-'a'] == 1 {
			return i
		}
	}
	return -1
}
