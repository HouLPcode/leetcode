/*
 * @lc app=leetcode id=914 lang=golang
 *
 * [914] X of a Kind in a Deck of Cards
 */
// 12 ms, faster than 95.16%
// [1,1,2,2,2,2]
// [1,1,1,1,2,2,2,2,2,2]
// 典型错误  不能用固定好的频次，需要时左右数字频次的最大公约数
func hasGroupsSizeX(deck []int) bool {
	// map 统计频次
	times := make(map[int]int)
	for _, v := range deck {
		times[v]++
	}
	X := 0 // 一定注意，X需要是所有频次的最大公约数
	for _, v := range times {
		X = v // 随机赋初值
		break
	}

	for _, v := range times {
		X = gcd(X, v)
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

func gcd(x, y int) int {
	if y%x == 0 {
		return x
	}
	return gcd(y%x, x)
}
