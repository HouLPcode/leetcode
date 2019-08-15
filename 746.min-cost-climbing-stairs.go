/*
 * @lc app=leetcode id=746 lang=golang
 *
 * [746] Min Cost Climbing Stairs
 */
// (4 ms) 92.67 %
// 典型dp
// dp[i] 表示跳到第i阶的最小花费 最终结果dp[l]
// dp[i] = min(dp[i-1]+num[i-1], dp[i-2]+num[i-2])
// 初始 dp[0] = 0, dp[1] = 0
func minCostClimbingStairs(cost []int) int {
	n := len(cost) // n [2,1000]
	pri1, pri2 := 0, 0
	for i := 2; i <= n; i++ {
		pri1, pri2 = pri2, min(pri1+cost[i-2], pri2+cost[i-1])
	}
	return pri2
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
