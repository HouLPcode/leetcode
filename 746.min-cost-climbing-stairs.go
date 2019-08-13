/*
 * @lc app=leetcode id=746 lang=golang
 *
 * [746] Min Cost Climbing Stairs
 */
// (4 ms) 92.57 %
// 典型dp
// dp[i] 表示跳到第i阶的最小花费 最终结果dp[0]
// dp[i] = cost[i]+ min(dp[i-1]+, dp[i-2])
// 初始 dp[0] = 0, dp[1] = 0
func minCostClimbingStairs(cost []int) int {
	dp0, dp1 := 0, 0
	rtn := 0
	for i := len(cost) - 1; i >= 0; i-- {
		rtn = cost[i] + min(dp0, dp1)
		dp0, dp1 = dp1, rtn
	}
	return min(dp0, dp1)
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
