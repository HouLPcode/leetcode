/*
 * @lc app=leetcode id=130 lang=golang
 *
 * [130] Surrounded Regions
 */
//  典型思路错误
//思路错误，DFS染色会有错误，导致不应该染色的点染色了 
// [["X","X","X","X"],
//  ["X","O","O","X"],
//  ["X","O","O","X"],
//  ["X","O","X","X"]]
func solve(board [][]byte)  {
	// 染色问题，BFS队列实现
	rows := len(board) // 不考虑空
	cols := len(board[0])

	visited := make([][]bool, rows)
	for i:=0; i<rows; i++ {
		visited[i] = make([]bool, cols)
	}

	for i:=0; i<rows; i++{
		for j:=0; j<cols; j++{
			if board[i][j] == 'O'{ //按行按列找到起始O的位置
				solveDFS(&board, i, j, rows, cols, &visited)
			}
		}
	}
}

var dx = []int{ 1, -1, 0,  0}
var dy = []int{ 0,  0, 1, -1}

func solveDFS(board *[][]byte, r,c,rows,cols int, visited *[][]bool){
	// 染色条件没有判断好
	(*visited)[r][c] = true
	// 先递归，再赋值
	for i:=0; i<4; i++ {
		if 0<=r+dx[i] && r+dx[i]<=rows-1 && 0<=c+dy[i] && c+dy[i]<=cols-1 && //不越界
			(*visited)[r+dx[i]][c+dy[i]]== false &&
			(*board)[r+dx[i]][c+dy[i]] == 'O' {
			solveDFS(board, r+dx[i], c+dy[i], rows, cols, visited)
		}
	}
	// 染色条件，不在边缘&周围未访问过的全是X
	for i:=0; i<4; i++ {
		if r+dx[i] < 0 || r+dx[i] > rows-1 || c+dy[i] < 0 || c+dy[i] > cols-1 || // 当前节点[r][c]不在边缘,而不是周围节点不在边缘，此处判断不越界
			((*visited)[r+dx[i]][c+dy[i]]== false && (*board)[r+dx[i]][c+dy[i]] == 'O') {// 未访问过的点都是X
			(*visited)[r][c] = false
			return // 直接返回，不赋值
		}
	}
	(*board)[r][c] = 'X'
	(*visited)[r][c] = false
}