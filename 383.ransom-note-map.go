/*
 * @lc app=leetcode id=383 lang=golang
 *
 * [383] Ransom Note
 */
//  map
func canConstruct(ransomNote string, magazine string) bool {
	alpMap := make(map[byte]int, len(magazine))
	for i := 0; i < len(magazine); i++ {
		alpMap[magazine[i]]++
	}
	for i := 0; i < len(ransomNote); i++ {
		if alpMap[ransomNote[i]] == 0 {
			return false
		} else {
			alpMap[ransomNote[i]]--
		}
	}
	return true
}
