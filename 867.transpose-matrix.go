/*
 * @lc app=leetcode id=867 lang=golang
 *
 * [867] Transpose Matrix
 */
//  24 ms 99.1 %
// 辅助空间
func transpose(A [][]int) [][]int {
	// 创建返回silce
	rows := len(A)
	if rows == 0 {
		return [][]int{}
	}
	cols := len(A[0])
	AT := make([][]int, cols)
	for c:=0; c < cols; c++ {
		AT[c] = make([]int, rows)
	}
	for r:=0; r<rows; r++ {
		for c:=0; c<cols; c++ {
			AT[c][r] = A[r][c]
		}
	}
	return AT
}

