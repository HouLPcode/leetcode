/*
 * @lc app=leetcode id=1160 lang=golang
 *
 * [1160] Find Words That Can Be Formed by Characters
 */
// (12 ms) √ Your runtime beats 100 %
func countCharacters(words []string, chars string) int {
	// 数组map
	charMap := make([]int, 26)
	for i := 0; i < len(chars); i++ {
		charMap[chars[i]-'a']++
	}

	tmp := make([]int, 26)
	sum := 0
	for i := 0; i < len(words); i++ {
		copy(tmp, charMap)
		sum += isGood(words[i], tmp)
	}
	return sum
}

func isGood(s string, chars []int) int {
	for i := 0; i < len(s); i++ {
		if chars[s[i]-'a'] == 0 {
			return 0
		}
		chars[s[i]-'a']--
	}
	return len(s)
}
