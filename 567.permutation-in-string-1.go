/*
 * @lc app=leetcode id=567 lang=golang
 *
 * [567] Permutation in String
 */
//  88 ms, faster than 26.14% 待优化 O(s1*s2)
func checkInclusion(s1 string, s2 string) bool {
	// 统计s1中每个字符的个数
	buf := make([]int, 26)
	for i := 0; i < len(s1); i++ {
		buf[int(s1[i]-'a')]++
	}
	tmp := make([]int, 26)
	for i := 0; i+len(s1)-1 < len(s2); {
		copy(tmp, buf)
		j := i
		for ; j < i+len(s1); j++ {
			tmp[int(s2[j]-'a')]--
			if tmp[int(s2[j]-'a')] < 0 {
				break
			}
		}
		if j == i+len(s1) {
			return true
		}
		i++ // 待优化，此处的i可以不从i++开始，跟j有关系？？？？？？？？？？？？
	}
	return false
}
