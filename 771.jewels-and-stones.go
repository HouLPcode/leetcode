/*
 * @lc app=leetcode id=771 lang=golang
 *
 * [771] Jewels and Stones
 */
//  0 ms
func numJewelsInStones(J string, S string) int {
	jmap := make(map[byte]bool, len(J))
	for i := 0; i < len(J); i++ {
		jmap[J[i]] = true
	}
	rtn := 0
	for i := 0; i < len(S); i++ {
		if jmap[S[i]] {
			rtn++
		}
	}
	return rtn
}

