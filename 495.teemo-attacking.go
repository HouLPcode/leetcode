/*
 * @lc app=leetcode id=495 lang=golang
 *
 * [495] Teemo Attacking
 */
//  48 ms, faster than 16.00%
func findPoisonedDuration(timeSeries []int, duration int) int {
	if len(timeSeries) == 0 { // 空数组的处理
		return 0
	}
	cnt := 0
	for i := 0; i < len(timeSeries)-1; i++ {
		if timeSeries[i+1]-timeSeries[i] >= duration {
			cnt += duration
		} else {
			cnt += (timeSeries[i+1] - timeSeries[i])
		}
	}
	return cnt + duration // 注意不要忘了最后一次眩晕时间
}
