/*
 * @lc app=leetcode id=1023 lang=golang
 *
 * [1023] Camelcase Matching
 */
//  (0 ms)
func camelMatch(queries []string, pattern string) []bool {
	rtn := make([]bool, len(queries))
	for i := 0; i < len(queries); i++ {
		rtn[i] = isParten(queries[i], pattern)
	}
	return rtn
}

// ，，
func isParten(query, pattern string) bool {
	p := 0
	q := 0
	for ; q < len(query); q++ { // 遍历query
		if query[q] == pattern[p] { // 和pattern相同 ，pattern右移
			p++
			if p == len(pattern) {
				break
			}
		} else if query[q] >= 'a' && query[q] <= 'z' { // 和pattern不同，则必须是小写字母
			// 小写字母
		} else {
			return false
		}
	}
	// pattern 有剩余
	if p != len(pattern) {
		return false
	}

	// 遍历完pattern后可能还有剩余
	q++
	for ; q < len(query); q++ { // 剩余的query必须是小写
		if query[q] < 'a' || query[q] > 'z' {
			return false
		}
	}
	return true
}
