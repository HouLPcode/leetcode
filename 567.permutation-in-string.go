/*
 * @lc app=leetcode id=567 lang=golang
 *
 * [567] Permutation in String
 */
//  (0 ms) √ Your runtime beats 100 %
// [l,r]维护一个窗口，比较窗口中字符出现的次数和s1中是否一致
func checkInclusion(s1 string, s2 string) bool {
	// 统计s1中每个字符的个数
	buf := make([]int, 26)
	for i := 0; i < len(s1); i++ {
		buf[int(s1[i]-'a')]++
	}
	tmp := make([]int, 26)
	left, right := 0, 0
	for ; right < len(s2); right++ {
		tmp[int(s2[right]-'a')]++
		if right-left+1 == len(s1) {
			// 遍历两个slice是都一样
			i := 0
			for ; i < 26; i++ { // 26步
				if buf[i] != tmp[i] {
					break
				}
			}
			if i == 26 {
				return true
			}
			tmp[int(s2[left]-'a')]--
			left++
		}
	}
	return false
}
