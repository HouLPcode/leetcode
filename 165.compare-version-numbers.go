/*
 * @lc app=leetcode id=165 lang=golang
 *
 * [165] Compare Version Numbers
 */
//  0 ms
func compareVersion(version1 string, version2 string) int {
	v1 := splitPoint(version1)
	v2 := splitPoint(version2)
	i := 0
	for ; i < len(v1) && i < len(v2); i++ {
		if v1[i] > v2[i] {
			return 1
		} else if v1[i] < v2[i] {
			return -1
		}
	}
	for ; i < len(v1); i++ {
		if v1[i] > 0 {
			return 1
		}
	}
	for ; i < len(v2); i++ {
		if v2[i] > 0 {
			return -1
		}
	}
	return 0
}

// 根据 . 切分string
func splitPoint(s string) []int {
	rtn := make([]int, 0, len(s))
	data := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '.' {
			rtn = append(rtn, data)
			data = 0
		} else {
			data = data*10 + int(s[i]-'0')
		}
	}
	// 一定注意，不要忘了最后一个数字，该数字不是由 . 切分
	rtn = append(rtn, data)
	return rtn
}
