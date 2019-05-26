/*
 * @lc app=leetcode id=76 lang=golang
 *
 * [76] Minimum Window Substring
 */
 func minWindow(s string, t string) string {
	minstr := ""
	minlen := len(s) // ?????
	// tMap := make(map[byte]bool, 128)
	// for i:=0; i<len(t); i++{
	// 	tMap[t[i]] = true
	// }
	for p1, p2 := 0, len(t)-1; p2 < len(s); p2++ {
		// T中一个也不包含，p1++
		if find(s[p1:p2+1], t) == false {
			p1++
		} else if findAll(s[p1:p2+1], t) { // 区间中包含全部T字母
			// 此时p2存储的是T中的某个字符
			if p2-p1+1 < minlen {
				minlen = p2 - p1 + 1
				minstr = s[p1 : p2+1]
			}
			p1++ //待优化，p1应该移动到第一个T中第二次出现的字母的位置
		}
	}
	return minstr
}

// ****************典型错误： 目标t中也可能有重复字符串*********
// bitmap 两个int64表示每个字符串中出现的字符，重复的计1次
func find(s, t string) bool { //包含t中某字母
	sb := str2bits(s)
	tb := str2bits(t)
	return sb[0]&tb[0] != 0 || sb[1]&tb[1] != 0
}

func findAll(s, t string) bool { //包含t中某字母
	sb := str2bits(s)
	tb := str2bits(t)
	return (sb[0]&tb[0] == tb[0]) && (sb[1]&tb[1] == tb[1])
}

func str2bits(s string) [2]uint64 {
	strbits := [2]uint64{0, 0}
	for i := 0; i < len(s); i++ {
		if s[i] < byte(64) {
			strbits[0] |= 1 << s[i]
		} else {
			strbits[1] |= 1 << (s[i] - 64)
		}
	}
	return strbits
}