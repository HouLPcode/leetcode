/*
 * @lc app=leetcode id=1010 lang=golang
 *
 * [1010] Pairs of Songs With Total Durations Divisible by 60
 */
// (24 ms) √ Your runtime beats 100 %
// [457,263,360]
func numPairsDivisibleBy60(time []int) int {
	// 统计 time%60=x 的个数，缓存后统计总数
	times := make([]int, 60)
	for _, v := range time {
		times[v%60]++
	}
	cnt := 0
	if times[0] > 1 {
		cnt = times[0] * (times[0] - 1) / 2 // 组合 C(x,2)
	}

	for i := 1; i < 30; i++ {
		cnt += times[i] * times[60-i]
	}

	if times[30] > 1 {
		cnt += times[30] * (times[30] - 1) / 2 // 组合 C(x,2)
	}
	return cnt
}
