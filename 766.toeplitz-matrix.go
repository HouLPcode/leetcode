/*
 * @lc app=leetcode id=766 lang=golang
 *
 * [766] Toeplitz Matrix
 *
 * https://leetcode.com/problems/toeplitz-matrix/description/
 */
//  4 ms 98.4 %
func isToeplitzMatrix(matrix [][]int) bool {
	rows := len(matrix)
	cols := len(matrix[0]) // 题目给出非空
	for r := rows - 1; r >= 0; r-- {
		if isValid(matrix, r, 0, rows, cols) == false {
			return false
		}
	}
	for c := 1; c < cols; c++ {
		if isValid(matrix, 0, c, rows, cols) == false {
			return false
		}
	}
	return true
}

func isValid(matrix [][]int, r, c, rows, cols int) bool {
	cur := matrix[r][c]
	for ; r < rows && c < cols; r, c = r+1, c+1 {
		if matrix[r][c] != cur {
			return false
		}
	}
	return true
}
