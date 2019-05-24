/*
 * @lc app=leetcode id=187 lang=golang
 *
 * [187] Repeated DNA Sequences
 */
func findRepeatedDnaSequences(s string) []string {
	// hash方式
	// s肯定满足条件，不会出现长度小于10的情况
	// if len(s) < 10{
	// 	return []string{}
	// }
	strMap := make(map[string]int, len(s)-9)
	result := []string{}
	for i:=0; i<len(s)-9; i++ {
		strMap[s[i:i+10]]++
	}
	for k,v := range strMap{
		if v > 1{
			result = append(result, k)
		}
	}
	return result
}

