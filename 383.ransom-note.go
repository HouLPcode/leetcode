/*
 * @lc app=leetcode id=383 lang=golang
 *
 * [383] Ransom Note
 */
//  4 ms  98.19%
//  通过[]slice替换map进行加速
func canConstruct(ransomNote string, magazine string) bool {
	alpMap := make([]int, 26)
	for i := 0; i < len(magazine); i++ {
		alpMap[magazine[i]-'a']++
	}
	for i := 0; i < len(ransomNote); i++ {
		if alpMap[ransomNote[i]-'a'] == 0 {
			return false
		} else {
			alpMap[ransomNote[i]-'a']--
		}
	}
	return true
}
