/*
 * @lc app=leetcode id=122 lang=golang
 *
 * [122] Best Time to Buy and Sell Stock II
 */
// (4 ms) √ Your runtime beats 91.99 %
// DP
// 买卖次数不限， dp[i][0/1] 0表示第i天没有股票的最大利润，1表示第i天有股票的最大利润
//  dp[i][0] = max(dp[i-1][0], dp[i-1][1] + nums[i]) // i-1没股票或者卖了
//  dp[i][1] = max(dp[i-1][0] - nums[i], dp[i-1][1])
// 最终结果  dp[l][0]  肯定是手里面没有股票的时候利润最多，即 dp[l][0]>dp[l][1]肯定成立
// 初始条件 dp[0][0] = 0 dp[0][1] = -nums[0]
func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	have, nohave := -prices[0], 0 // 初始条件
	for i := 1; i < len(prices); i++ {
		have, nohave = max(nohave-prices[i], have), max(nohave, have+prices[i])
	}

	return nohave
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
