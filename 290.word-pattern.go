/*
 * @lc app=leetcode id=290 lang=golang
 *
 * [290] Word Pattern
 */
 // "aba"\n"cat cat cat dog"
//  "abba"\n"dog dog dog dog"
func wordPattern(pattern string, str string) bool {
    // str 按空格分词后，跟pattern中的一个字符配对，str或pattern替换成另一个，然后匹配结果
	strs := strings.Split(str, " ")
	alpMap := make(map[string]byte, 0)
	for k,v := range strs{
		if k >= len(pattern){
			return false
		}
		if _,ok := alpMap[v]; !ok { // v没出现过，但是pattern出现，导致结果错误
			alpMap[v] = pattern[k] // alpMap["dog"] = 'a'
		}else if alpMap[v] == pattern[k]{
			continue
		}else{
			return false
		}
	}
	return true	
	// tmpStr := ""
	// for i:=0; i<len(strs); i++{
	// 	tmpStr += alpMap[pattern[i]] + " "//注意删除最后的空格
	// }
	// return tmpStr == str+" "
}

