/*
 * @lc app=leetcode id=796 lang=golang
 *
 * [796] Rotate String
 */
func rotateString(A string, B string) bool {
	// 暴力旋转每种可能
	if len(A) <=1 && A == B{ // 不用旋转
		return true
	}
	for i:=1; i<len(A); i++{ // 旋转n-1次
		if A[i:]+A[:i] == B{
			return true
		} 
	}
	return false
}

