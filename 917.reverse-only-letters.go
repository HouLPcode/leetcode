/*
 * @lc app=leetcode id=917 lang=golang
 *
 * [917] Reverse Only Letters
 */
// 0 ms
func reverseOnlyLetters(S string) string {
	// 新创建空间，双向遍历
	l := len(S)
	rtn := make([]byte, 0, l)
	for i, j := 0, l-1; i < l; {
		if isLetter(S[i]) == false {
			rtn = append(rtn, S[i])
			i++
		} else if isLetter(S[j]) == false {
			j--
		} else { // i是字母，j是字母
			rtn = append(rtn, S[j])
			j--
			i++
		}
	}
	return string(rtn)
}

func isLetter(a byte) bool {
	return (a >= 'a' && a <= 'z') || (a >= 'A' && a <= 'Z')
}
