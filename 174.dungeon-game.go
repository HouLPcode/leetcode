/*
 * @lc app=leetcode id=174 lang=golang
 *
 * [174] Dungeon Game
 */
// [[100]] 
func calculateMinimumHP(dungeon [][]int) int {
	// dp[i][j] 右下到[i][j]需要的最少血量
	// dp[i][j] = min(dp[i][j+1], dp[i+1][j]) - nums[i][j] && dp[i][j] >= 1
	// 返回dp[0][0]
	rows := len(dungeon)
	cols := len(dungeon[0]) // 不用考虑越界问题
	dp := make([][]int, rows)
	for k := 0; k<rows; k++{
		dp[k] =  make([]int, cols)
	}

	// 先处理最后一行一列
	dp[rows-1][cols-1] = -dungeon[rows-1][cols-1] + 1 // 注意只有最后一个+1，其他的不需要+1
	if dp[rows-1][cols-1] < 1{
		dp[rows-1][cols-1] = 1
	}
	for r:=rows-2; r>=0; r--{
		dp[r][cols-1] = dp[r+1][cols-1] - dungeon[r][cols-1]
		if dp[r][cols-1] < 1{
			dp[r][cols-1] = 1
		}
	} 
	for c:=cols-2; c>=0; c--{
		dp[rows-1][c] = dp[rows-1][c+1] - dungeon[rows-1][c] 
		if dp[rows-1][c] < 1{
			dp[rows-1][c] = 1
		}
	}
	for r:=rows-2; r>=0; r--{
		for c:=cols-2; c>=0; c--{
			dp[r][c] = min(dp[r+1][c], dp[r][c+1]) -dungeon[r][c]
			if dp[r][c] < 1{
				dp[r][c] = 1
			}
		}
	}
	return dp[0][0]
}

func min(a, b int) int{
	if a < b{
		return a
	}
	return b
}
