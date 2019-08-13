/*
 * @lc app=leetcode id=1025 lang=golang
 *
 * [1025] Divisor Game
 */
// (0 ms)
// 法1 ： dp怎么解决？？？
// 法2 ： 数学归纳法 偶数true 奇数false
func divisorGame(N int) bool {
	return N&1 == 0
}
