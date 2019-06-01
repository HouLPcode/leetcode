/*
 * @lc app=leetcode id=796 lang=golang
 *
 * [796] Rotate String
 */
func rotateString(A string, B string) bool {
	// 法1：暴力旋转每种可能
	// 法2：在B+B中找A
	if len(A) != len(B){ // 长度不同直接返回false
		return false
	}
	return strings.Contains(B+B, A)
}

