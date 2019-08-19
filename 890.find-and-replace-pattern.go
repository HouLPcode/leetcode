/*
 * @lc app=leetcode id=890 lang=golang
 *
 * [890] Find and Replace Pattern
 */
// (0 ms) √ Your runtime beats 100 %
// 注意s p映射是一对一，需要判断一对多，多对一时返回false
func findAndReplacePattern(words []string, pattern string) []string {
	rtn := []string{}
	for i := 0; i < len(words); i++ {
		if isValid(words[i], pattern) {
			rtn = append(rtn, words[i])
		}
	}
	return rtn
}

func isValid(s, p string) bool {
	// 只有小写字母
	maps := make([]byte, 26)
	visitedp := int32(0)
	if len(s) != len(p) {
		return false
	}
	for i := 0; i < len(s); i++ {
		if maps[s[i]-'a'] != byte(0) { // 出现过
			if maps[s[i]-'a'] != p[i] { // 映射不一样了，一对多
				return false
			}
		} else { // 新的映射
			if visitedp&(1<<(p[i]-'a')) != 0 { // p出现过，多对一
				return false
			}
			maps[s[i]-'a'] = p[i]
			visitedp |= (1 << (p[i] - 'a'))
		}
	}
	return true
}
