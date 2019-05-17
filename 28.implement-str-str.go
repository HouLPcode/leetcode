/*
 * @lc app=leetcode id=28 lang=golang
 *
 * [28] Implement strStr()
 */
 // "aaa"\n"aaaa" 注意needle在haystack末尾时的情况
 // KMP算法？？？
func strStr(haystack string, needle string) int {
    if len(needle) == 0{
		return 0
	}
	for i:=0; i<len(haystack); i++{
		if haystack[i] == needle[0]{
			j:=1
			for ; j<len(needle)&&i+j<len(haystack)&&haystack[i+j] == needle[j];j++{}
			if j == len(needle){
				return i
			}
		}
	}
	return -1
}

