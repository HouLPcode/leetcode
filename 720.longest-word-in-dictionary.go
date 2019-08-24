/*
 * @lc app=leetcode id=720 lang=golang
 *
 * [720] Longest Word in Dictionary
 */
 // 16 ms, faster than 77.55%
 // 方法1 ： trie
 // 方法2： sort + map 以下是该方法实现
func longestWord(words []string) string {
	sort.Strings(words) // 排序后保证前缀都在word前面，且多个结果中字母序
	wordMap := make(map[string]bool, 0)
	rtn := ""
	for _,v := range words{
		if len(v) == 1 || wordMap[v[:len(v)-1]] == true{
			// 单词只有一个字符，或者前缀存在
			if len(v) > len(rtn) {
				rtn = v
			}
			wordMap[v] = true
		}
	}
	return rtn
}

