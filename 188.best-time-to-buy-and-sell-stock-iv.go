/*
 * @lc app=leetcode id=188 lang=golang
 *
 * [188] Best Time to Buy and Sell Stock IV
 */
// 最多买卖k次。 不限次数的情况下最多买卖 len/2 次
//  dp[i][j][0] 第i天最多买j次手里没股票的利润 ... 最多j次还是恰好j次？？？
//  dp[i][j][1] 第i天最多买j次手里有股票的利润
// dp[i][j][0] = max(dp[i-1][j][0], dp[i-1][j][1]+nums[i])
// dp[i][j][1] = max(dp[i-1][j][1], dp[i-1][j-1][0]-nums[i])
// 初始 dp[0][1][1] = -nums[0]， 其他情况 dp[0][.][.] = 0
func maxProfit(k int, prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	dp := make([][][2]int, len(prices))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([][2]int, k+1) // k+1
	}

	dp[0][1][1] = -prices[0]

	for i := 1; i < len(prices); i++ { // i=1表示第二天
		// k = 0, dp = 0
		// i天最多买 i/2+1 次
		// for j := 1; j <= k && j < i/2+1; j++ {
		for j := 1; j <= k; j++ {
			dp[i][j][0], dp[i][j][1] = max(dp[i-1][j][0], dp[i-1][j][1]+prices[i]), max(dp[i-1][j][1], dp[i-1][j-1][0]-prices[i])
		}
	}
	// return max(dp[len(prices)-1][k][0], dp[len(prices)-1][len(prices)/2][0])
	return dp[len(prices)-1][k][0]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
