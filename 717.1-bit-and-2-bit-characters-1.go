/*
 * @lc app=leetcode id=717 lang=golang
 *
 * [717] 1-bit and 2-bit Characters
 */
//  4 ms, faster than 33.33%
func isOneBitCharacter(bits []int) bool {
	if len(bits) == 1 { //0
		return true
	}
	if bits[len(bits)-2] == 0 { // ....00
		return true
	}
	if bits[len(bits)-2] == 1 && (len(bits) == 2 || bits[len(bits)-3] == 0) { // ...010 或 10
		return false
	}

	// 有1 必须携带后面一位 ...110
	i := 0
	for i < len(bits)-1 {
		if bits[i] == 1 {
			i += 2 //忽略1后面的一位
		} else {
			i++
		}
	}
	if i == len(bits)-1 {
		return true
	}
	return false
}
