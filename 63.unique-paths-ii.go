/*
 * @lc app=leetcode id=63 lang=golang
 *
 * [63] Unique Paths II
 */
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	buf := [100][100]int{}//题目中给出最大值为100*100
	row := len(obstacleGrid)
	col := len(obstacleGrid[0])//不考虑越界
	
	if obstacleGrid[0][0] == 0{//要考虑第一块就是障碍吗？
		buf[0][0] = 1
	} 
	// 初始化边界值
	for i:=1;i<row;i++{
		if obstacleGrid[i][0] == 0{
			buf[i][0] = buf[i-1][0]
		} 
	}
	for j:=1;j<col;j++{
		if obstacleGrid[0][j] == 0{
			buf[0][j] = buf[0][j-1]
		} 
	}

	for i:=1;i<row;i++{
		for j:=1;j<col;j++{
			if obstacleGrid[i][j] == 0{
				buf[i][j] = buf[i-1][j] + buf[i][j-1]
			}
		}
	}

	return buf[row-1][col-1]
}

