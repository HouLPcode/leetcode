/*
 * @lc app=leetcode id=801 lang=golang
 *
 * [801] Minimum Swaps To Make Sequences Increasing
 */
//  8 ms, faster than 93.55%
// DP 解决
// dp[i][0] 表示第i位没有交换，达到严格递增的最小交换次数
// dp[i][1] 表示第i位交换，达到严格递增的最小交换次数
// 转移方程
// 取最小值
// dp[i][0] = dp[i-1][0] // i-1 不交换， i不交换
// dp[i][0] = dp[i-1][1] // i-1  交换， i不交换
// 取最小值
// dp[i][1] = dp[i-1][0]+1 // i-1 不交换， i交换
// dp[i][1] = dp[i-1][1]+1 // i-1 交换， i交换
// 最终结果 min(dp[i][0,1])
// 初始条件 dp[0] = {0, 1}
func minSwap(A []int, B []int) int {
	pri := [2]int{0, 1}
	cur := [2]int{100, 100}
	for i := 1; i < len(A); i++ {
		// 一定注意循环进来的时候将cur赋值成非常大的值，避免不进第一个if带来的影响
		cur[0] = len(A)
		cur[1] = len(A)
		if A[i] > A[i-1] && B[i] > B[i-1] { // 都不交换，或者都交换
			cur[0] = pri[0]
			cur[1] = pri[1] + 1
		}
		if A[i] > B[i-1] && B[i] > A[i-1] { // 交换1个
			cur[0] = min(cur[0], pri[1])
			cur[1] = min(cur[1], pri[0]+1)
		}
		pri = cur
	}
	return min(cur[0], cur[1])
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
