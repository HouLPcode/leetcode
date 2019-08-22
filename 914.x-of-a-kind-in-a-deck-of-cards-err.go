/*
 * @lc app=leetcode id=914 lang=golang
 *
 * [914] X of a Kind in a Deck of Cards
 */
// [1,1,2,2,2,2]
// [1,1,1,1,2,2,2,2,2,2]
// 典型错误  不能用固定好的频次，需要时左右数字频次的最小公约数
func hasGroupsSizeX(deck []int) bool {
	// map 统计频次
	times := make(map[int]int)
	for _, v := range deck {
		times[v]++
	}
	X := 10000
	for _, v := range times { // 不能随便取出来一个数字，需要是最小的数字
		if X > v {
			X = v
		}
	}
	if X < 2 {
		return false
	}
	for _, v := range times {
		if v != X && v%X != 0 { // 注意，此处%表示不能分成多组
			return false
		}
	}
	return true
}
