/*
 * @lc app=leetcode id=539 lang=golang
 *
 * [539] Minimum Time Difference
 */
// 0 ms
// ["01:01","02:01"]
//  24 * 60 = 1440 类似桶排序
// 先转换成分钟插入桶中，然后白能力桶
func findMinDifference(timePoints []string) int {
	mark := make([]bool, 1440)
	for i := 0; i < len(timePoints); i++ {
		mins := int(timePoints[i][0]-'0')*600 + int(timePoints[i][1]-'0')*60 +
			int(timePoints[i][3]-'0')*10 + int(timePoints[i][4]-'0')
		if mark[mins] {
			return 0 // 有重复时间
		}
		mark[mins] = true
	}
	sub := 1440 // MAX
	pre := -1
	num1 := -1 // 第一个时间
	for i := 0; i < 1440; i++ {
		if mark[i] == false { // 忽略没有出现过的时间
			continue
		}
		if pre == -1 {
			pre = i
			num1 = i
		} else {
			sub = min(sub, i-pre)
			pre = i
		}
	}
	return min(sub, num1+1440-pre) // 首尾时间差00:00-23:59 = 1
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
