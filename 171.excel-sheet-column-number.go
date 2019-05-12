/*
 * @lc app=leetcode id=171 lang=golang
 *
 * [171] Excel Sheet Column Number
 */
func titleToNumber(s string) int {
	//可以理解成26进制编码
	rnt := 0
	for _,v := range s{ // 全是大写字母，不需要转为rune
		rnt = rnt * 26 + int(v - 'A') + 1
	}
	return rnt
}

