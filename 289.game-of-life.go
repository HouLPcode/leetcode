/*
 * @lc app=leetcode id=289 lang=golang
 *
 * [289] Game of Life
 */
//  (0 ms)
// 原地变换，增加新的元素表示状态
// 一次遍历记录更改状态
// 二次遍历更新数值
//  活 <2  活->死    2
//  活 2/3 活->活
//  活 >3  活->死    2
//  死 =3  死->活    3

var dirs = [][]int{
	[]int{-1, -1},
	[]int{-1, 0},
	[]int{-1, 1},
	[]int{0, -1},
	[]int{0, 1},
	[]int{1, -1},
	[]int{1, 0},
	[]int{1, 1},
}

func gameOfLife(board [][]int) {
	rows := len(board)
	if rows == 0 {
		return
	}
	cols := len(board[0])

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			cnt := 0
			for i := 0; i < 8; i++ {
				if r+dirs[i][0] >= 0 && r+dirs[i][0] < rows && c+dirs[i][1] >= 0 && c+dirs[i][1] < cols &&
					(board[r+dirs[i][0]][c+dirs[i][1]] == 1 || board[r+dirs[i][0]][c+dirs[i][1]] == 2) {
					cnt++
				}
			}
			if board[r][c] == 0 {
				if cnt == 3 {
					board[r][c] = 3
				}
			} else {
				switch {
				case cnt < 2:
					board[r][c] = 2
				case cnt == 2 || cnt == 3:
					// board[r][c] = 0
				case cnt > 3:
					board[r][c] = 2
				}
			}
		}
	}
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if board[r][c] == 2 {
				board[r][c] = 0
			} else if board[r][c] == 3 {
				board[r][c] = 1
			}
		}
	}
}
