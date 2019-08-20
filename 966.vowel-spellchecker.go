/*
 * @lc app=leetcode id=966 lang=golang
 *
 * [966] Vowel Spellchecker
 */
// 876 ms, faster than 16.67% 待优化
//   × testcase: '["wg","uo","as","kv","ra","mw","gi","we","og","zu"]\n["ov"]'
func spellchecker(wordlist []string, queries []string) []string {
	rtn := []string{}
	for i := 0; i < len(queries); i++ {
		level := 4
		str := 0
		for j := 0; j < len(wordlist); j++ {
			l := check(wordlist[j], queries[i])
			if level > l {
				level = l
				str = j
			}
		}
		if level == 3 {
			rtn = append(rtn, "")
		} else {
			rtn = append(rtn, wordlist[str])
		}
	}
	return rtn
}

// 一摸一样是0
// 大小写 1
// 元音 2
// 3
func check(word, query string) int {
	if len(word) != len(query) {
		return 3
	}
	rtn := 0
	for i := 0; i < len(word); i++ {
		if word[i] == query[i] {

		} else if word[i] == query[i]+'A'-'a' || word[i] == query[i]-'A'+'a' {
			if rtn < 1 {
				rtn = 1
			}
		} else if (query[i] == 'a' || query[i] == 'e' || query[i] == 'i' || query[i] == 'o' || query[i] == 'u' ||
			query[i] == 'A' || query[i] == 'E' || query[i] == 'I' || query[i] == 'O' || query[i] == 'U') &&
			(word[i] == 'a' || word[i] == 'e' || word[i] == 'i' || word[i] == 'o' || word[i] == 'u' ||
				word[i] == 'A' || word[i] == 'E' || word[i] == 'I' || word[i] == 'O' || word[i] == 'U') { // 注意：元音只能替换元音
			if rtn < 2 {
				rtn = 2
			}
		} else {
			return 3
		}
	}
	return rtn
}
