/*
 * @lc app=leetcode id=168 lang=golang
 *
 * [168] Excel Sheet Column Title
 */
func convertToTitle(n int) string {
	// 26进制,一共26个字符，不是27进制
	// 字母与数字运算转字符串
	str := "" 
	for n != 0{ //从末尾向头转换
		str = string(byte('A') + byte((n-1)%26)) + str // i=26
		n = (n-1)/26 // 注意是对n-1进行的运算
	}
	return str
}

