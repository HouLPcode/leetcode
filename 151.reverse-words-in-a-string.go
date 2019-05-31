/*
 * @lc app=leetcode id=151 lang=golang
 *
 * [151] Reverse Words in a String
 */
//  " " ""
func reverseWords(s string) string {
	strs := strings.Fields(s)
	if len(strs) == 0 { // 注意，此处判断的是分词后的长度，避免"" " "等的影响
		return ""
	}
	str := ""
	for i:=len(strs)-1; i>=0; i-- {
		str += strs[i] + " "
	}
	return str[:len(str)-1] // 去掉最后的空格，空的时候此处会导致越界，注意添加判断
}

