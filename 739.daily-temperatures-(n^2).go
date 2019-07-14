/*
 * @lc app=leetcode id=739 lang=golang
 *
 * [739] Daily Temperatures
 */
//  1444 ms  15.64%
func dailyTemperatures(T []int) []int {
	// 暴力破解，O(n^2)
	l := len(T)
	rtn := make([]int, l)
	for i, v := range T {
		for j := i + 1; j < l; j++ {
			if T[j] > v {
				rtn[i] = j - i
				break
			}
		}
	}
	return rtn
}
