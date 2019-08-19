/*
 * @lc app=leetcode id=916 lang=golang
 *
 * [916] Word Subsets
 */
// (552 ms) √ Your runtime beats 82.61 %
func wordSubsets(A []string, B []string) []string {
	// 统计B中每个单词中字符出现的次数，选出最大的值
	cntB := make([]int, 26)

	for i := 0; i < len(B); i++ {
		tmp := make([]int, 26)
		for j := 0; j < len(B[i]); j++ {
			tmp[B[i][j]-'a']++
		}
		for j := 0; j < 26; j++ {
			if tmp[j] > cntB[j] {
				cntB[j] = tmp[j]
			}
		}
	}

	rtn := []string{}
	for i := 0; i < len(A); i++ {
		tmp := make([]int, 26)
		copy(tmp, cntB)
		for j := 0; j < len(A[i]); j++ {
			if tmp[A[i][j]-'a'] > 0 { // B 中有这个字符，计数--
				tmp[A[i][j]-'a']--
			}
		}
		k := 0
		for ; k < 26; k++ { // 查看B的字符有没有剩余
			if tmp[k] > 0 {
				break
			}
		}
		if k == 26 {
			rtn = append(rtn, A[i])
		}
	}
	return rtn
}
