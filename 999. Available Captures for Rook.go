// 0 ms
func numRookCaptures(board [][]byte) int {
	// 8*8
	cnt := 0
	r, c := 0, 0
	for ; r < 8; r++ {
		for c = 0; c < 8; c++ {
			if board[r][c] == 'R' {
				break
			}
		}
		if c < 8 {
			break
		}
	}
	// r , c
	// 行
	for c1 := c + 1; c1 < 8; c1++ {
		if board[r][c1] == '.' {

		} else if board[r][c1] == 'p' {
			cnt++
			break
		} else {
			break
		}
	}
	for c1 := c - 1; c1 >= 0; c1-- {
		if board[r][c1] == '.' {

		} else if board[r][c1] == 'p' {
			cnt++
			break
		} else {
			break
		}
	}
	// 列
	for r1 := r + 1; r1 < 8; r1++ {
		if board[r1][c] == '.' {

		} else if board[r1][c] == 'p' {
			cnt++
			break
		} else {
			break
		}
	}
	for r1 := r - 1; r1 >= 0; r1-- {
		if board[r1][c] == '.' {

		} else if board[r1][c] == 'p' {
			cnt++
			break
		} else {
			break
		}
	}
	return cnt
}
