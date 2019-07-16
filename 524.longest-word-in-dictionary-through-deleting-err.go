import "strings"

/*
 * @lc app=leetcode id=524 lang=golang
 *
 * [524] Longest Word in Dictionary through Deleting
 */
//  "aewfafwafjlwajflwajflwafj"
// ["apple","ewaf","awefawfwaf","awef","awefe","ewafeffewafewf"]
// 审题错误，典型错误，不能只通过计数判断是否满足条件，还需要注意字母顺序
func findLongestWord(s string, d []string) string {
	// 通过map缓存字母个数，逐个比较
	sbuf := make([]int, 26)
	for i := 0; i < len(s); i++ {
		sbuf[int(s[i]-'a')]++
	}
	rtn := ""
	for i := 0; i < len(d); i++ {
		databuf := make([]int, 26)
		copy(databuf, sbuf)
		if isValid(d[i], databuf) == true {
			if len(rtn) < len(d[i]) {
				rtn = d[i]
			} else if len(rtn) == len(d[i]) && strings.Compare(rtn, d[i]) > 0 {
				rtn = d[i] // 返回值需要按照字母表顺序优先返回
			}
		}
	}
	return rtn
}

func isValid(target string, buf []int) bool {
	for i := 0; i < len(target); i++ {
		buf[int(target[i]-'a')]--
		if buf[int(target[i]-'a')] < 0 {
			return false
		}
	}
	return true
}
