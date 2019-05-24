/*
 * @lc app=leetcode id=290 lang=golang
 *
 * [290] Word Pattern
 */
 // "aba"\n"cat cat cat dog"
//  "abba"\n"dog dog dog dog"
// "abba"\n"dog cat cat fish"
// "jquery"\n"jquery"
func wordPattern(pattern string, str string) bool {
    // str 按空格分词后，跟pattern中的一个字符配对，str或pattern替换成另一个，然后匹配结果
	strs := strings.Split(str, " ")
	if len(strs) != len(pattern){
		return false
	}
	alpMap := make(map[string]byte, 0)
	visit := [128]bool{}// 一定注意，需要统计pattern中的字符是否出现过
	for k,v := range strs{
		if _,ok := alpMap[v]; !ok && visit[int(pattern[k])]==false { // v没出现过，但是pattern出现，导致结果错误
			alpMap[v] = pattern[k] // alpMap["dog"] = 'a'
			visit[int(pattern[k])] = true
		}else if alpMap[v] == pattern[k]{
			continue
		}else{
			return false
		}
	}
	return true	
}

