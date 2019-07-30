/*
 * @lc app=leetcode id=242 lang=golang
 *
 * [242] Valid Anagram
 */
// 0 ms
func isAnagram(s string, t string) bool {
	m1 := make([]int, 26) //一共有26个小写字母
	m2 := make([]int, 26)
	for i := 0; i < len(s); i++ {
		m1[s[i]-'a']++
	}
	for i := 0; i < len(t); i++ {
		m2[t[i]-'a']++
	}
	for i := 0; i < 26; i++ {
		if m1[i] != m2[i] {
			return false
		}
	}
	return true
}
