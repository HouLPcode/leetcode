/*
 * @lc app=leetcode id=258 lang=golang
 *
 * [258] Add Digits
 */
func addDigits(num int) int {
	//规律
	// 0 -> 0
	// 1...9 -> 1..9
	// 10...18 -> 1...9
	// 19...27 -> 1...9
	if num == 0{
		return 0
	}
	if num%9 == 0{
		return 9
	}
	return num%9
}

