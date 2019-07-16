/*
 * @lc app=leetcode id=345 lang=golang
 *
 * [345] Reverse Vowels of a String
 */
//  4 ms, faster than 91.46%
func reverseVowels(s string) string {
	// 两个指针左右开始遍历，遇到元音就反转
	l, r := 0, len(s)-1
	rtn := []byte(s)
	lenth := len(s)

	for ; l < lenth && isVowel(rtn[l]) == false; l++ {
	}
	for ; r >= 0 && isVowel(rtn[r]) == false; r-- {
	}
	for l < r {
		rtn[l], rtn[r] = rtn[r], rtn[l]
		l++
		r--
		for ; l < lenth && isVowel(rtn[l]) == false; l++ {
		}
		for ; r >= 0 && isVowel(rtn[r]) == false; r-- {
		}
	}
	return string(rtn)
}

func isVowel(a byte) bool { // 注意原因包括大小写
	return a == 'a' || a == 'e' || a == 'i' || a == 'o' || a == 'u' ||
		a == 'A' || a == 'E' || a == 'I' || a == 'O' || a == 'U'
}
