/*
 * @lc app=leetcode id=205 lang=golang
 *
 * [205] Isomorphic Strings
 */
// 0 ms
//  "ab"\n"aa"
func isIsomorphic(s string, t string) bool {
	convMap := make(map[byte]byte, len(s))
	tused := make(map[byte]bool, len(s))
	// if len(s) != len(t) { // 题目给出长度相同
	// 	return false // 防止长度不同导致下面循环访问越界
	// }
	for i := 0; i < len(s); i++ {
		if v, ok := convMap[s[i]]; !ok { // 典型错误，不能只看s是否访问过，s的赋值呀要看t访问过没
			if tused[t[i]] == true {
				return false
			}
			convMap[s[i]] = t[i]
			tused[t[i]] = true
		} else if v == t[i] {
			// 访问过，s -> t对应关系没变
		} else {
			return false //  访问过，s -> t对应关系变化，直接返回错误
		}
	}
	return true
}
