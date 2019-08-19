/*
 * @lc app=leetcode id=680 lang=golang
 *
 * [680] Valid Palindrome II
 */
// (20 ms) √ Your runtime beats 85 %
// 分别尝试错过左右字符一次，判断是不是回文
func validPalindrome(s string) bool {
	l, r := 0, len(s)-1
	for l <= r && s[l] == s[r] {
		l++
		r--
	}

	if l > r { // 本身就是回文
		return true
	}

	// 过滤掉s[l]
	l1 := l + 1
	r1 := r
	for l1 <= r1 && s[l1] == s[r1] {
		l1++
		r1--
	}
	if l1 > r1 { // 本身就是回文
		return true
	}
	// 过滤掉 s[r]
	l2 := l
	r2 := r - 1
	for l2 <= r2 && s[l2] == s[r2] {
		l2++
		r2--
	}
	if l2 > r2 { // 本身就是回文
		return true
	}
	return false
}
