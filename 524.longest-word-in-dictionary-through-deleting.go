import "strings"

/*
 * @lc app=leetcode id=524 lang=golang
 *
 * [524] Longest Word in Dictionary through Deleting
 */
//  20 ms, faster than 77.46%
//  "aewfafwafjlwajflwajflwafj"
// ["apple","ewaf","awefawfwaf","awef","awefe","ewafeffewafewf"]
func findLongestWord(s string, d []string) string {
	rtn := ""
	for i := 0; i < len(d); i++ {
		if isValid(s, d[i]) == true {
			if len(rtn) < len(d[i]) {
				rtn = d[i]
			} else if len(rtn) == len(d[i]) && strings.Compare(rtn, d[i]) > 0 {
				rtn = d[i] // 返回值需要按照字母表顺序优先返回
			}
		}
	}
	return rtn
}

func isValid(tmp, target string) bool {
	i, j := 0, 0
	for i < len(tmp) && j < len(target) {
		if tmp[i] == target[j] {
			j++
		}
		i++
	}
	return j == len(target)
}
