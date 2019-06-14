/*
 * @lc app=leetcode id=292 lang=golang
 *
 * [292] Nim Game
 */
func canWinNim(n int) bool {
	// 规律，能被4整除的数，肯定输
	// return n%4 != 0
	return n&0x3 != 0 //bit操作占用内存更少
}

