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
	if len(strs) != len(pattern){//注意长度不匹配时直接报错
		return false
	}
	alpMap := make(map[byte]string, 0)
	for k,v := range strs{
		// k 有越界的可能
		// 错误："ab"\n"dog dog"
		alpMap[pattern[k]] = v  // alpMap['a'] = "dog"
	}
	tmpStr := ""
	for i:=0; i<len(pattern); i++{
		tmpStr += alpMap[pattern[i]] + " "//注意删除最后的空格
	}
	return tmpStr == str+" "
}

