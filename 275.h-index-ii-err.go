/*
 * @lc app=leetcode id=275 lang=golang
 *
 * [275] H-Index II
 */
//   × testcase: '[0,0]'
// [0]
// 被二分法卡住了，不会写
func hIndex(citations []int) int {
	l := len(citations)
	if l == 0 {
		return 0
	}

	// 折半尝试??????????????????
	s, e := 0, l-1
	for s <= e {
		mid := (e-s)/2 + s
		if citations[mid] < l-mid+1 { // 引用数 < 文章数
			s = mid + 1
		} else {
			e = mid - 1
		}
	}
	if e == 0 {
		return l
	}
	return l - e // 文章数是索引+1
}
