/*
 * @lc app=leetcode id=14 lang=golang
 *
 * [14] Longest Common Prefix
 */
func longestCommonPrefix(strs []string) string {
	// 法1. 两个比较，直到最后一个
	// 法2. 一个字符一个字符比较
	// 以下两个比较
	if len(strs) == 0{
		return ""
	}
	str := strs[0]
	for i:=1; i<len(strs);i++{
		str = prefix(str,strs[i])
	}
	return str
}

func prefix(s1,s2 string) string{
	str := ""
	for i:=0; i<len(s1)&&i<len(s2);i++{
		if s1[i] == s2[i]{
			str += string(s1[i])
		}else{
			break
		}
	}
	return str
}

