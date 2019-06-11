/*
 * @lc app=leetcode id=657 lang=golang
 *
 * [657] Robot Return to Origin
 */
// func judgeCircle(moves string) bool {
// 	// buf 缓存
// 	bufmap := make(map[byte]int, 4)
// 	for i:=0; i<len(moves); i++ {
// 		bufmap[moves[i]]++
// 	}
// 	return bufmap['U'] == bufmap['D'] && bufmap['L'] == bufmap['R']
// }
func judgeCircle(moves string) bool {
	u,d,l,r := 0,0,0,0
	for i:=0; i<len(moves); i++ {
		switch moves[i]{
		case 'U': u++
		case 'D': d++
		case 'L': l++
		case 'R': r++
		}
	}
	return u == d && l == r
}
