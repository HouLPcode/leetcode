/*
 * @lc app=leetcode id=64 lang=golang
 *
 * [64] Minimum Path Sum
 */
func minPathSum(grid [][]int) int {
    if grid == nil || len(grid) == 0{
		return 0
	}
	
	// f[i][j] 到第[i][j] 的路径和的最小值
	// f[i][j] = min( f[i-1][j], f[i][j-1]) + grid[i][j]
	// 结果 f[i][j]
	buf := [][]int{}// TODO 减少空间

	rows := len(grid)
	cols := len(grid[0])

	for i:=0; i<rows; i++{
		buf = append(buf,make([]int,cols,cols))
	}
	buf[0][0] = grid[0][0]
	for i:=1; i<rows; i++{
		buf[i][0] = buf[i-1][0] + grid[i][0]
	}
	for j:=1; j<cols; j++{
		buf[0][j] = buf[0][j-1] + grid[0][j]
	}

	for i:=1; i<rows; i++{
		for j:=1; j<cols; j++{
			buf[i][j] = min(buf[i-1][j],buf[i][j-1]) + grid[i][j]
		}
	}
	return buf[rows-1][cols-1]
}

func min(a,b int)int{
	if a < b{
		return a
	}
	return b
}
