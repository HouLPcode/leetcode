/*
 * @lc app=leetcode id=200 lang=golang
 *
 * [200] Number of Islands
 */
//  0 ms
 // 法1 DFS染色法实现如下,会改变grid入参内容
// 法2 并查集
func numIslands(grid [][]byte) int { 
	rows := len(grid)
	if rows == 0 {
		return 0
	}
	cols := len(grid[0])
	rtn := 0
	for r:=0; r < rows; r++ {
		for c:=0; c < cols; c++ {
			if grid[r][c] == '1' {
				rtn++ //新的岛屿,计数++
				dfs(grid, r, c, rows, cols) // 染色周围的'1'
			}
		}
	}
	return rtn
}

var dirs = [][]int{
	[]int{1, 0},
	[]int{0, 1},
	[]int{-1, 0},
	[]int{0, -1},
}

func dfs(grid [][]byte, r, c, rows, cols int) {
	if r < 0 || r == rows || c < 0 || c == cols {
		return 
	}
	if grid[r][c] == '0' {
		return 
	}
	grid[r][c] = '0' // 染色
	for i:=0; i<4; i++ {
		dfs(grid, r+dirs[i][0], c+dirs[i][1], rows, cols)
	}
}
