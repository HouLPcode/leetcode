/*
 * @lc app=leetcode id=709 lang=golang
 *
 * [709] To Lower Case
 */
//  0 ms
func toLowerCase(str string) string {
	rtn := []byte(str)
	for i:=0; i<len(rtn); i++ {
		if rtn[i] >= 'A' && rtn[i] <= 'Z' {
			rtn[i] = rtn[i] - 'A' + 'a'
		}
	}
	return string(rtn)
}

