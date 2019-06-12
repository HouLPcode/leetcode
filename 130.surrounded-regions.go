/*
 * @lc app=leetcode id=130 lang=golang
 *
 * [130] Surrounded Regions
 */
//  28 ms  99.22%
//思路错误，DFS染色会有错误，导致不应该染色的点染色了 [["X","X","X","X"],["X","O","O","X"],["X","O","O","X"],["X","O","X","X"]]
// 思路转换，先把不符合要求的点标记成别的元素，然后再恢复回来
// 1.首先不由左上到右下找第一个O元素，改成从四周找O，并DFS将其变成第三个元素*，
// 2.然后遍历所有元素，将O -> X
// 3.最后将所有 * -> O,流程结束
func solve(board [][]byte)  {
	// 染色问题，BFS队列实现
	rows := len(board) //
	if rows == 0 {
		return 
	}
	cols := len(board[0])

	// 遍历边缘元素
	for i:=0; i<rows; i++{
		if board[i][cols-1] == 'O'{ // O->*
			solveDFS(&board, i, cols-1, rows, cols)
		}
		if board[i][0] == 'O'{ // O->*
			solveDFS(&board, i, 0, rows, cols)
		}
	}
	for j:=0; j<cols; j++{
		if board[rows-1][j] == 'O'{ // O->*
			solveDFS(&board, rows-1, j, rows, cols)
		}
		if board[0][j] == 'O'{ // O->*
			solveDFS(&board, 0, j, rows, cols)
		}
	}
	// O -> X
	for i:=0; i<rows; i++{
		for j:=0; j<cols; j++{
			if board[i][j] == 'O'{ //按行按列找到起始O的位置
				board[i][j] = 'X'
			}
		}
	}
	// *->O
	for i:=0; i<rows; i++{
		for j:=0; j<cols; j++{
			if board[i][j] == '*'{ //按行按列找到起始O的位置
				board[i][j] = 'O'
			}
		}
	}
}

var dx = []int{ 1, -1, 0,  0}
var dy = []int{ 0,  0, 1, -1}

func solveDFS(board *[][]byte, r,c,rows,cols int){
	(*board)[r][c] = '*'
	for i:=0; i<4; i++ {
		if 0<=r+dx[i] && r+dx[i]<=rows-1 && 0<=c+dy[i] && c+dy[i]<=cols-1 && //不越界
		   (*board)[r+dx[i]][c+dy[i]] == 'O' {
			solveDFS(board, r+dx[i], c+dy[i], rows, cols)
		}
	}
}


