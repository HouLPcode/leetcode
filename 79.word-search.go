/*
 * @lc app=leetcode id=79 lang=golang
 *
 * [79] Word Search
 */
//  回溯法：类似DFS下的组合
// 4 ms 98.82 %
func exist(board [][]byte, word string) bool {
	rows := len(board)
	if rows == 0 {
		return false
	}
	cols := len(board[0])
	for r:=0; r<rows; r++ {
		for c:=0; c<cols; c++ {
			if dfs(board, word, 0, r, c, rows, cols) == true {
				return true
			}
		}
	}
	return false
}
var dirs = [][]int{
	[]int{1, 0},
	[]int{0, 1},
	[]int{-1, 0},
	[]int{0, -1},
}
//board是candidates，r，c构成起始位置，word[s]为target值, 
// used标志，中间过程值，最终结果集合 忽略
func dfs(candidates [][]byte,word string, s, r, c, rows, cols int) bool {
	if r<0 || r == rows || c < 0 || c == cols {
		return false //边缘越界
	}
	if candidates[r][c] != word[s] {
		return false // 不是要找的值，返回false
	}
	if s == len(word) - 1 {
		return true // 单词寻找完成
	}

	// 小技巧：通过将当前访问值candidate[r][c]赋值为target中不存在的值，节省used空间
	cur := candidates[r][c]
	candidates[r][c] = '0'
	for i:=0; i<4; i++ {
		if dfs(candidates, word, s+1, r+dirs[i][0], c+dirs[i][1], rows, cols) {
			return true
		}
	}
	candidates[r][c] = cur
	return false
	// defer func(r, c int){
	// 	candidates[r][c] = cur
	// }(r, c)
	// return dfs(candidates, word, s+1, r+1, c, rows, cols) ||
	// 		dfs(candidates, word, s+1, r, c+1, rows, cols) ||
	// 		dfs(candidates, word, s+1, r-1, c, rows, cols) ||
	// 		dfs(candidates, word, s+1, r, c-1, rows, cols)
}

