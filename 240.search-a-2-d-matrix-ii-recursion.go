/*
 * @lc app=leetcode id=240 lang=golang
 *
 * [240] Search a 2D Matrix II
 */
//  24 ms 93.81%
func searchMatrix(matrix [][]int, target int) bool {
	// 分治法尝试, 从中间节点分
	rows := len(matrix)
	if rows == 0 {
		return false
	}
	cols := len(matrix[0])
	return search(matrix, 0, rows-1, 0, cols-1, target)
}

func search(matrix [][]int, sr, er, sc, ec, target int) bool {
	if sr > er || sc > ec {
		return false
	}
	midr := (er-sr)/2 + sr
	midc := (ec-sc)/2 + sc
	if matrix[midr][midc] == target {
		return true
	} else if matrix[midr][midc] > target {
		// 排除右下
		return search(matrix, sr, midr-1, sc, midc, target) || //左上 ...
			search(matrix, sr, midr-1, midc+1, ec, target) || //右上 ..
			search(matrix, midr, er, sc, midc-1, target) // 左下 ...
	} else {
		// 排除左上
		return search(matrix, sr, midr-1, midc+1, ec, target) || //右上 ...
			search(matrix, midr+1, er, sc, midc, target) || // 左下...
			search(matrix, midr, er, midc+1, ec, target) // 右下 ...
	}
}
