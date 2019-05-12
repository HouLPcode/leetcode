/*
 * @lc app=leetcode id=387 lang=golang
 *
 * [387] First Unique Character in a String
 */
func firstUniqChar(s string) int {
	// 第一遍缓存每个字符出现的次数
	// 第二遍找到频次为1的索引并返回
	// 时间O(n)
	buf := make(map[rune]int)
	for _,v := range s{
		buf[v]++
	}
	for k,v := range s{
		if buf[v] == 1{
			return k
		}
	}
	return -1// 没有则返回 -1
}

