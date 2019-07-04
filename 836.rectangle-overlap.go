/*
 * @lc app=leetcode id=836 lang=golang
 *
 * [836] Rectangle Overlap
 */
//  0 ms
// 类似直线的的区间重合判断条配件，以不重合为思考点，不满足则重合
func isRectangleOverlap(rec1 []int, rec2 []int) bool {
	// 不重合的条件好判断
	// 忽略不合法输入
	if rec1[2] <= rec2[0] || rec1[0] >= rec2[2] || // 左右
		rec1[3] <= rec2[1] || rec1[1] >= rec2[3] { // 上下
		return false
	}
	return true
}

