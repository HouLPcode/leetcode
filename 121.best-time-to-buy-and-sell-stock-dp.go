/*
 * @lc app=leetcode id=121 lang=golang
 *
 * [121] Best Time to Buy and Sell Stock
 */
// (4 ms) √ Your runtime beats 96.48 %
// [7,6,4,3,1]
// DP 尝试
// 只能买卖一次股票，
// dp[i][0][0] 表示第i天的最大利润， [][0][] 没买过，  [][1][] 买过， [][][0]手里没有， [][][1]手里有股票

// dp[i][0][0] = 0
// dp[i][0][1] = 不可能
// dp[i][1][0] = max(dp[i-1][1][0], dp[i-1][1][1]+nums[i])
// dp[i][1][1] = max(dp[i-1][1][1], dp[i-1][0][0]-nums[i]) = max(dp[i-1][1][1], -nums[i])
// 最终结果  dp[l][1][0]
// 初始条件 dp[0][1][0]=min  dp[0][1][1]=-nums[0]
func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	dp_1_0, dp_1_1 := 0, -prices[0] // 初始条件,一定注意dp_1_0不能初始化为math.MinInt64，因为最终利润不可能为负数，利润最小为0
	for i := 1; i < len(prices); i++ {
		dp_1_0, dp_1_1 = max(dp_1_0, dp_1_1+prices[i]), max(dp_1_1, -prices[i])
	}

	return dp_1_0
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
