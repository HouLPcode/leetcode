/*
 * @lc app=leetcode id=290 lang=golang
 *
 * [290] Word Pattern
 */
func wordPattern(pattern string, str string) bool {
    // str 按空格分词后，跟pattern中的一个字符配对，str或pattern替换成另一个，然后匹配结果
	strs := strings.Split(str, " ")
	alpMap := make(map[byte]string, 0)
	for k,v := range strs{
		alpMap[pattern[k]] = v  // alpMap['a'] = "dog"
	}
	tmpStr := ""
	for i:=0; i<len(pattern); i++{
		tmpStr += alpMap[pattern[i]] + " "//注意删除最后的空格
	}
	return tmpStr == str+" "
}

