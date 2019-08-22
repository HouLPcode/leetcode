/*
 * @lc app=leetcode id=123 lang=golang
 *
 * [123] Best Time to Buy and Sell Stock III
 */
// err--------------------------------
// [2,1,2,0,1]
// [1,2,3,4,5]
//最多可以完成 两笔 交易
//  dp[i][j][0] 第i天买j次手里没股票的利润
//  dp[i][j][1] 第i天买j次手里有股票的利润
// dp[][0][] = 0
// dp[i][1][0] = max(dp[i-1][1][0], dp[i-1][1][1]+prices[i])
// dp[i][1][1] = max(dp[i-1][1][1], dp[i-1][0][0]-prices[i]) = max(dp[i-1][1][1], -prices[i])
// dp[i][2][0] = max(dp[i-1][2][0], dp[i-1][2][1]+prices[i])
// dp[i][2][1] = max(dp[i-1][2][1], dp[i-1][1][0]-prices[i])
// 初始 dp[0][1][1] = -prices[0]， 其他情况 dp[0][.][.] = 0
// 最终结果 max(dp[l][1][0], dp[l][2][0])
func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	dp_1_0, dp_1_1, dp_2_0, dp_2_1 := 0, -prices[0], 0, 0

	for i := 2; i < len(prices); i++ {
		dp_1_0, dp_1_1, dp_2_0, dp_2_1 = max(dp_1_0, dp_1_1+prices[i]), max(dp_1_1, -prices[i]), max(dp_2_0, dp_2_1+prices[i]), max(dp_2_1, dp_1_0-prices[i])
	}
	return max(dp_1_0, dp_2_0)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
