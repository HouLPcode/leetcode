/*
 * @lc app=leetcode id=125 lang=golang
 *
 * [125] Valid Palindrome
 */
 // 测试集 "0P" 注意测试集中还有数字
 func isPalindrome(s string) bool {
	// 前后两个指针，逐步比较
	// 怎么忽略标点符号？？？？？？
	i, j := 0, len(s)-1
	for i < j{
		if isAlphanumeric(s[i]) == false {
			i++
		}else if isAlphanumeric(s[j]) == false {
			j--
		}else if s[i] == s[j] ||
			((s[i]<'0'|| s[i]>'9') && int(s[i]-s[j]) % int('a' - 'A')  == 0) {//大小写不敏感,注意此处 '0'-'P' == 'a'- 'A'
			i++ //有可能导致i>j
			j--
		} else {
			break
		}
	}
	if i < j{
		return false
	}
	return true
}

func isAlphanumeric(s byte) bool {
	if ('0' <= s && s <= '9') ||
		('a' <= s && s <= 'z') ||
		('A' <= s && s <= 'Z'){
		return true
	}
	return false
}
