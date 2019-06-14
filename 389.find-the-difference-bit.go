/*
 * @lc app=leetcode id=389 lang=golang
 *
 * [389] Find the Difference
 */
func findTheDifference(s string, t string) byte {
	// bit
	// byte类型底层是uint8，支持位运算，算数运算
	rtn := byte(0)
	for i:=0; i<len(s); i++ {
		rtn ^= s[i]
	}
	for i:=0; i<len(t); i++ {
		rtn ^= t[i]
	}
	return rtn
}

