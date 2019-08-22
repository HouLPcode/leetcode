/*
 * @lc app=leetcode id=717 lang=golang
 *
 * [717] 1-bit and 2-bit Characters
 */
//  (0 ms)
// 比进行末尾情况判断还快，不需要过度优化
func isOneBitCharacter(bits []int) bool {
	i := 0
	for i < len(bits) {
		if i == len(bits)-1 {
			return true
		}
		if bits[i] == 1 {
			i += 2 //忽略1后面的一位
		} else {
			i++
		}
	}
	return false
}
