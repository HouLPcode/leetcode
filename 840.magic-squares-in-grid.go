/*
 * @lc app=leetcode id=840 lang=golang
 *
 * [840] Magic Squares In Grid
 */
// [[5,5,5],[5,5,5],[5,5,5]]
// 暴力求解所有情况
func numMagicSquaresInside(grid [][]int) int {
	cnt := 0
	rows, cols := len(grid), len(grid[0])
	for r := 0; r+2 < rows; r++ {
		for c := 0; c+2 < cols; c++ {
			if isMagic(grid, r, c) {
				cnt++
			}
		}
	}
	return cnt
}

func isMagic(grid [][]int, sr, sc int) bool { // 3*3
	sum := 15 // 肯定是15

	tmp := 0
	for i := 0; i < 3; i++ { // row
		tmp = 0
		for j := 0; j < 3; j++ {
			tmp += grid[sr+i][sc+j]
		}
		if tmp != sum {
			return false
		}
	}
	for i := 0; i < 3; i++ { // col
		tmp = 0
		for j := 0; j < 3; j++ {
			tmp += grid[sr+j][sc+i]
		}
		if tmp != sum {
			return false
		}
	}

	if sum != grid[sr][sc]+grid[sr+1][sc+1]+grid[sr+2][sc+2] {
		return false
	}

	if sum != grid[sr][sc+2]+grid[sr+1][sc+1]+grid[sr+2][sc] {
		return false
	}
	return true
}
