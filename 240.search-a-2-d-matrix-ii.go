/*
 * @lc app=leetcode id=240 lang=golang
 *
 * [240] Search a 2D Matrix II
 */
//  24ms 96.6%
func searchMatrix(matrix [][]int, target int) bool {
	// 右上 -> 左下
	rows := len(matrix)
	if rows == 0 {
		return false
	}
	cols := len(matrix[0])

	for r,c:=0,cols-1; r<rows && c >= 0; {
		if matrix[r][c] == target{
			return true
		}else if matrix[r][c] < target {
			r++
		}else if matrix[r][c] > target {
			c--
		}
	}
	return false
}

