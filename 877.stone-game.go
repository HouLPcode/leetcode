/*
 * @lc app=leetcode id=877 lang=golang
 *
 * [877] Stone Game
 */
// 8 ms, faster than 14.00%
// 优化1：dp[i][j]先手+，后手-，最后和0比较
// 优化2：二维改为一维滚动数组

// 法1 ： 数学归纳法返回答案
// 法2： dp，下面进行解答. 博弈类问题
// dp[i][j][0,1] 表示在[i,j]范围内的先手最大和后手最大
// 转移方程：
// dp[i][j][0] = max( num[i] + dp[i+1][j][1], num[j] + dp[i][j-1][1] )
//     先手                 拿左     后手         拿右       后手
// dp[i][j][1] = max( dp[i+1][j][0],     dp[i][j-1][0])
//    后手           丢左，后手变先手       丢右，后手边先手
// 临界条件
// dp[i][i][0] = num[i] 只有一个元素，先手得到
// dp[i][i][1] = 0 只有一个元素，后手得0
// 最终结果 dp[0][len-1] [0]>[1] // 先手>后手 返回ture

// 时间复杂度 O(n*2) 空间 O(n*2)
func stoneGame(piles []int) bool {
	dp := make([][][2]int, len(piles))
	for i := 0; i < len(piles); i++ {
		dp[i] = make([][2]int, len(piles))
	}

	// 初始化base 条件
	for i := 0; i > len(piles); i++ {
		dp[i][i][0] = piles[i]
		dp[i][i][1] = 0
	}

	// 正方形，斜着遍历数组
	for i := 2; i <= len(piles); i++ {
		for row := 0; row <= len(piles)-i; row++ {
			col := row + i - 1

			dp[row][col][0] = max(piles[row]+dp[row+1][col][1], piles[col]+dp[row][col-1][1])
			dp[row][col][1] = max(dp[row+1][col][0], dp[row][col-1][0])
		}
	}
	return dp[0][len(piles)-1][0] > dp[0][len(piles)-1][1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
