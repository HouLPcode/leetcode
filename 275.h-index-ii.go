/*
 * @lc app=leetcode id=275 lang=golang
 *
 * [275] H-Index II
 */
//  (16 ms) √ Your runtime beats 88.46 %
//   × testcase: '[0,0]'
// [0]
// 被二分法卡住了，不会写
func hIndex(citations []int) int {
	l := len(citations)
	if l == 0 {
		return 0
	}

	// 逐步遍历，此处可以进行二分法优化
	i := l - 1
	for i >= 0 {
		if citations[i] < l-i { // 引用数 < 文章数
			break
		}
		i--
	}
	// if i == -1 {
	// 	return l
	// } else if i == l-1{
	// 	return 0
	// }
	return l - i - 1
}
