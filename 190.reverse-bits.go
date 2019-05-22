/*
 * @lc app=leetcode id=190 lang=golang
 *
 * [190] Reverse Bits
 */
func reverseBits(num uint32) uint32 {
	// 没有技巧，位运算交换bit位置即可
	// 通过折半方式减少一步一步运算的步骤
	num = (num >> 16) | (num << 16)
	num = (num&0xff00ff00 >> 8) | (num&0x00ff00ff << 8)
	num = (num&0xf0f0f0f0 >> 4) | (num&0x0f0f0f0f << 4)
	num = (num&0xcccccccc >> 2) | (num&0x33333333 << 2)
	num = (num&0xaaaaaaaa >> 1) | (num&0x55555555 << 1)
	return num
}

