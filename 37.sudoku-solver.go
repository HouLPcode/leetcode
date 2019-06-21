/*
 * @lc app=leetcode id=37 lang=golang
 *
 * [37] Sudoku Solver
 */
//  0 ms
//  DFS,回溯法
// r，c怎么转换到box小格子中 第[r/3*3+c/3]格
func solveSudoku(board [][]byte)  {
    rowbuf := make([][]bool, 9)
    colbuf := make([][]bool, 9)
	boxbuf := make([][]bool, 9)
	for i:=0; i<9; i++ {
		rowbuf[i] = make([]bool, 10)// 1-9是标志位，0位没有用，为了方便编程
		colbuf[i] = make([]bool, 10)
		boxbuf[i] = make([]bool, 10)
	}
	// 遍历一次，添加缓存
	for r:=0; r<9; r++ {
		for c:=0; c<9; c++ {
			if board[r][c] != '.' {
				num := board[r][c]-'0'
				rowbuf[r][num] = true
				colbuf[c][num] = true
				boxbuf[r/3*3+c/3][num] = true // r,c 转换到box方格中
			}
		}
	}

	// 不需要逐个位置遍历，从第一个开始，依次填写直到最后即可完成
	// for r:=0; r<9; r++ {
	// 	for c:=0; c<9; c++ {

	// 	}
	// } 
	dfs(board, 0, 0, &rowbuf, &colbuf, &boxbuf)
}
// candidate， s ， target， used， curr, ans 
// board是候选项， r,c是起始s， 没有target，...buf是used，curr和ans直接在board中体现
 func dfs(board [][]byte, r, c int, rowbuf, colbuf, boxbuf *[][]bool) bool {
	if r == 9 { //按行填写，直到[8][8]后面的出现，表示所有格子已经填写完成
		return true // 一行或一列填写完成
	}
	if board[r][c] != '.' {
		// TODO
		if c == 8 {
			return dfs(board, r+1, 0, rowbuf, colbuf, boxbuf) //一行一行的填写
		}
		return dfs(board, r, c+1, rowbuf, colbuf, boxbuf) //一行一行的填写
	}

	for i:=1; i<=9; i++ { // 候选项
		if (*rowbuf)[r][i]==true || (*colbuf)[c][i]==true || (*boxbuf)[r/3*3+c/3][i]==true {
			continue
		}
		(*rowbuf)[r][i]=true
		(*colbuf)[c][i]=true
		(*boxbuf)[r/3*3+c/3][i]=true
		board[r][c] = byte(i) + '0' 
		if c == 8 {
			if dfs(board, r+1, 0, rowbuf, colbuf, boxbuf) == true {//一行一行的填写
				return true
			}
		}else {
			if dfs(board, r, c+1, rowbuf, colbuf, boxbuf) == true {//一行一行的填写
				return true
			}
		}
		(*rowbuf)[r][i]=false
		(*colbuf)[c][i]=false
		(*boxbuf)[r/3*3+c/3][i]=false
		board[r][c] = '.' 
	}
	return false
}

