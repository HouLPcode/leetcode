/*
 * @lc app=leetcode id=76 lang=golang
 *
 * [76] Minimum Window Substring
 */
//  双指针和map缓存
// 1. map中缓存t总每个字符出现的次数
// 2. 双指针p1 p2组成窗口，统计窗口中每个字母出现的顺序，且比较s是否包含t
// 3. 当p1字符不在t中或数量大于t中数量的时候，尽可能的右移p1
// "a"\n"b"
// "ab"\n"A"
// "a"\n"a"
// "bba"\n"ab"
 func minWindow(s string, t string) string {
	minstr := ""
	minlen := len(s) // ?????
	tMap := make(map[byte]int, 128)
	for i:=0; i<len(t); i++{
		tMap[t[i]]++
	}
	windowMap := make(map[byte]int, 0)
	for p1, p2 := 0, 0; p2 < len(s); p2++ {
		windowMap[s[p2]]++
		// 一定注意此处的p1<p2的条件，保证下面语句不越界s[p1]，s[p1:p2+1]
		for p1 < p2 && tMap[s[p1]]==0 { //p1不是t中字符
			windowMap[s[p1]]--
			p1++
		}
		for p1 < p2 && windowMap[s[p1]] > tMap[s[p1]] { // p1数量多
			windowMap[s[p1]]--
			p1++
		} 
		if findAll(windowMap, tMap) == true{ // 区间中包含全部T字母
			// 此时p2存储的是T中的某个字符
			if p2-p1+1 <= minlen {// 注意此处是<=,保证初始更新 "a" "a"的情况
				minlen = p2 - p1 + 1
				minstr = s[p1 : p2+1]
			}
			windowMap[s[p1]]--
			p1++ //待优化，p1应该移动到T中第二次出现的字母的位置
		}
	}
	return minstr
}

func findAll(smap, tmap map[byte]int) bool { //包含t中某字母
	for k,v := range tmap{
		if smap[k] < v{
			return false
		}
	}
	return true
}
