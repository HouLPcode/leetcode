/*
 * @lc app=leetcode id=931 lang=golang
 *
 * [931] Minimum Falling Path Sum
 */
//  (12 ms) √ Your runtime beats 46.88 %
// dp[i][j] 下降到最底层i位置的最小路径和
// dp[i][j] = num[i][j] + min(dp[i-1][j-1], dp[i-1][j+1], dp[i-1][j])
// 输出结果 min(dp[i][0...l])
// 初始 dp[0][...] = num[0][...]
func minFallingPathSum(A [][]int) int {
	rows, cols := len(A), len(A[0])
	dp := make([][]int, rows)
	for i := 0; i < rows; i++ {
		dp[i] = make([]int, cols)
	}

	for i := 0; i < cols; i++ {
		dp[0][i] = A[0][i]
	}

	for r := 1; r < rows; r++ {
		dp[r][0] = min(dp[r-1][0], dp[r-1][1]) + A[r][0]
		dp[r][cols-1] = min(dp[r-1][cols-1], dp[r-1][cols-2]) + A[r][cols-1]

		for c := 1; c < cols-1; c++ {
			dp[r][c] = A[r][c] + min3(dp[r-1][c-1], dp[r-1][c], dp[r-1][c+1])
		}
	}
	rtn := dp[rows-1][0]
	for c := 0; c < cols; c++ {
		if dp[rows-1][c] < rtn {
			rtn = dp[rows-1][c]
		}
	}
	return rtn
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func min3(a, b, c int) int {
	if a <= b && a <= c {
		return a
	} else if b <= a && b <= c {
		return b
	}
	return c
}
