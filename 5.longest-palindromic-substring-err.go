/*
 * @lc app=leetcode id=5 lang=golang
 *
 * [5] Longest Palindromic Substring
 */
 // dp *****典型错误******* 当前的回文不一定是前面的回文+最后一个字符，可能是左边也扩展一个字符
func longestPalindrome(s string) string {
	// 找到字符串中的最长回文子串
	// df 尝试
	// f(i) 包含第i字符的最长回文  ?????不包含i字符是不是可以直接递归本函数
	// f(i) = s[i] or f(i-1)+1  *****典型错误******* 当前的回文不一定是前面的回文+最后一个字符，可能是左边也扩展一个字符
	// 最终结果 max[f(0,1,2...n)]
	rtn := ""
	for i:=1; i<=len(s); i++{//多次重复计算，待优化
		str := palindrome(s[:i])
		if len(str) > len(rtn){
			rtn = str
		}
	}
	return rtn
}
func palindrome(s string) string{
	if len(s) == 0 || len(s) == 1{
		return s
	}
	str := palindrome(s[:len(s)-1])//f(i-1)
	if str==s[len(s)-len(str)-1:len(s)-1] && isPalindrome(str+s[len(s)-1:]){
		//f(i-1)的回文结合当前字符串
		return s[len(s)-len(str)-1:]
	}
	return s[len(s)-1:]//不是回文，则直接返回当前字符
}

func isPalindrome(s string) bool{
	i,j := 0,len(s)-1
	for i < j{
		if s[i] != s[j]{
			return false
		}
		i++
		j--
	}
	return true
}

