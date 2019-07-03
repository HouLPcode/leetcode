import "strings"

/*
 * @lc app=leetcode id=824 lang=golang
 *
 * [824] Goat Latin
 */
//  0 ms
//  "Each word consists of lowercase and uppercase letters only"
func toGoatLatin(S string) string {
	words := strings.Split(S, " ") // []string
	for k, v := range words {
		word := []byte(v)
		// 注意此处是大小写字母都是
		if word[0] == 'a' || word[0] == 'e' || word[0] == 'i' || word[0] == 'o' || word[0] == 'u' ||
			word[0] == 'A' || word[0] == 'E' || word[0] == 'I' || word[0] == 'O' || word[0] == 'U' {

		} else {
			word = append(word, word[0])
			word = word[1:]
		}
		word = append(word, []byte{'m', 'a'}...)
		for i := 0; i < k+1; i++ {
			word = append(word, 'a')
		}
		words[k] = string(word)
	}
	rtn := ""
	for _, v := range words {
		rtn += " "
		rtn += v
	}
	return rtn[1:]
}

