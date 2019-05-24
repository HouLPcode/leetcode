/*
 * @lc app=leetcode id=409 lang=golang
 *
 * [409] Longest Palindrome
 */
func longestPalindrome(s string) int {
	// 统计每个字母出现的次数，
	// 最大长度即为所有偶数+所有（奇数-1）+ 1个剩余元素
	alp := make(map[byte]int, 0)
	for i:=0; i<len(s); i++ {
		alp[s[i]]++
	}
	cnt := 0
	mflag := false // 标记是否有奇数剩余
	for _,v := range alp{
		if v & 1 == 0{
			cnt += v
		}else{
			cnt += v -1
			mflag = true
		}
	}
	if mflag{
		cnt++
	}
	return cnt
}

