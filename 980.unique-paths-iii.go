/*
 * @lc app=leetcode id=980 lang=golang
 *
 * [980] Unique Paths III
 */
//  0 ms
 // DFS 回溯法尝试
//  注意，除了-1，剩余的所有点都需要访问到
func uniquePathsIII(grid [][]int) int {
	rows := len(grid)
	if rows == 0 {
		return 0
	}
	cols := len(grid[0])
	ans := 0
	for r:=0; r<rows; r++ {
		for c:=0; c<cols; c++ {
			if grid[r][c] == 1 {
				dfs(grid, r, c, rows, cols, &ans)
				// break // 跳出几层循环
				return ans
			}
		}
	}
	return ans
}

var dirs = [][]int{
	[]int{-1, 0},
	[]int{0, 1},
	[]int{1, 0},
	[]int{0, -1},
}

func dfs(grid [][]int, r, c, rows, cols int, ans *int) {
	if r < 0 || r == rows || c < 0 || c == cols {
		return
	}
	if grid[r][c] == -1 {
		return
	}
	if grid[r][c] == 2 { //找到目标值
		for r:=0; r<rows; r++ { // 除了-1，剩余的所有点都需要访问到
			for c:=0; c<cols; c++ {
				if grid[r][c] == 0 {
					return // 还有未访问的点
				}
			}
		}
		(*ans)++
		return
	}
	
	grid[r][c] = -1 // 标记为不可访问
	for i:=0; i<4; i++ {
		dfs(grid, r+dirs[i][0], c+dirs[i][1], rows, cols, ans) 
	}
	grid[r][c] = 0 // 恢复可访问标记
}

