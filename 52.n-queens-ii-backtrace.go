/*
 * @lc app=leetcode id=52 lang=golang
 *
 * [52] N-Queens II
 */
// 0 ms
func totalNQueens(n int) int {
	rtn := 0
	chessboard := make([][]bool, n)
	for i := 0; i < n; i++ {
		chessboard[i] = make([]bool, n)
	}
	backtrace_queen(&chessboard, 0, n, &rtn)
	return rtn
}

// 不记录攻击范围，不同queen的攻击范围有重叠，没法回退
func backtrace_queen(chessboard *[][]bool, r, n int, count *int) {
	// 按行放置queen，到最后一行后表示找到了一个解
	if r == n {
		(*count) += 1
		return
	}
	// 在下一行中遍历放置queen
	for c := 0; c < n; c++ {
		if isValid(*chessboard, r, c, n) == false {
			continue // 在攻击范围内
		}
		(*chessboard)[r][c] = true
		backtrace_queen(chessboard, r+1, n, count)
		(*chessboard)[r][c] = false
	}
}

func isValid(chessboard [][]bool, r, c, n int) bool {
	// 查看当前位置是否放置queen
	for i := 0; i < n; i++ { // row
		if chessboard[r][i] == true {
			return false // 在攻击范围
		}
	}
	for i := 0; i < n; i++ { // col
		if chessboard[i][c] == true {
			return false // 在攻击范围
		}
	}
	for i, j := r+1, c+1; i < n && j < n; i, j = i+1, j+1 { // 右下
		if chessboard[i][j] == true {
			return false // 在攻击范围
		}
	}
	for i, j := r-1, c-1; i >= 0 && j >= 0; i, j = i-1, j-1 { // 左上
		if chessboard[i][j] == true {
			return false // 在攻击范围
		}
	}
	for i, j := r+1, c-1; i < n && j >= 0; i, j = i+1, j-1 { // 左下
		if chessboard[i][j] == true {
			return false // 在攻击范围
		}
	}
	for i, j := r-1, c+1; i >= 0 && j < n; i, j = i-1, j+1 { // 右上
		if chessboard[i][j] == true {
			return false // 在攻击范围
		}
	}
	return true
}
