/*
 * @lc app=leetcode.cn id=1124 lang=golang
 *
 * [1124] 表现良好的最长时间段
 */
// 典型错误，思路错误，可以右前缀0
// [6,9,9]
func longestWPI(hours []int) int {
	// 统计窗口中的劳累-不劳累天数，负值直接舍去
	rtn := 0
	curMax := 0
	time := 0 // 劳累-不劳累
	for i := 0; i < len(hours); i++ {
		if hours[i] > 8 {
			time++
		} else {
			time--
		}
		if time > 0 {
			curMax++
			if curMax > rtn {
				rtn = curMax
			}
		} else {
			time = 0
			curMax = 0
		}
	}
	if curMax > rtn {
		rtn = curMax
	}
	return rtn
}
