/*
 * @lc app=leetcode id=744 lang=golang
 *
 * [744] Find Smallest Letter Greater Than Target
 */
//  0 ms
func nextGreatestLetter(letters []byte, target byte) byte {
	sub := byte(26) // 小写字母的差值 1-26,注意，相同字母差值是26，
	rtn := byte(0)
	for _, v := range letters {
		if v == target { // 题目要求letters肯定有至少2个不同数字，所以与target相同字母永远不可能是返回值
			continue
		} else if (v-target+26)%26 < sub {
			sub = (v - target + 26) % 26
			rtn = v
		}
	}
	return rtn
}

