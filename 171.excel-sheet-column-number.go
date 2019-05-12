/*
 * @lc app=leetcode id=171 lang=golang
 *
 * [171] Excel Sheet Column Number
 */
func titleToNumber(s string) int {
	//可以理解成26进制编码
	rnt := 0
	for _,v := range s{
		rnt = rnt * 26 + int(v - 'A') + 1
	}
	return rnt
}

