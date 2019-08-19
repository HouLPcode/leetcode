/*
 * @lc app=leetcode id=791 lang=golang
 *
 * [791] Custom Sort String
 */
//  (0 ms)
func customSortString(S string, T string) string {
	// 从右往左统计T中S元素的个数
	// 顺便把不在S中的字符移位到最右侧
	cntS := make([]int, 26)
	for i := 0; i < len(S); i++ {
		cntS[S[i]-'a']++ // 此处多统计一个
	}
	Ts := []byte(T)
	p := len(Ts) - 1 // p右侧的全是不在S中的字符,不包含p
	for i := len(Ts) - 1; i >= 0; i-- {
		if cntS[Ts[i]-'a'] > 0 { // 该元素在S中出现
			cntS[Ts[i]-'a']++
		} else { // 没有出现过
			Ts[i], Ts[p] = Ts[p], Ts[i]
			p--
		}
	}
	// [0:p+1] 是S中出现的字符
	j := 0
	for i := 0; i < len(S); i++ {
		for cntS[S[i]-'a'] > 1 { // 此处记得是1，多统计了一个
			cntS[S[i]-'a']--
			Ts[j] = S[i]
			j++
		}
	}

	return string(Ts)
}
