// 0 ms, faster than 100.00%
// DP
// DP[i][j] i个骰子总和为j的组合数
// dp[i][j] = dp[i-1][j-1] + dp[i-1][j-2] + ... dp[i-1][j-f]
// 结果：dp[d][target]
// 初始 dp[1][1...f] = 1,  i个骰子总数范围[i,f*i]
func numRollsToTarget(d int, f int, target int) int {
	if target < d || target > f*d {
		return 0
	}

	// d*target 的数组
	buf := make([][]int, d+1)
	for i := 0; i < d+1; i++ {
		buf[i] = make([]int, target+1)
	}

	buf[0][0] = 1

	for i := 1; i <= d; i++ { // 骰子数
		for j := 1; j <= f; j++ { // 数值
			for k := j; k <= target; k++ { //注意此处的填充方式
				buf[i][k] = (buf[i][k] + buf[i-1][k-j]) % 1000000007
			}
		}
	}
	return buf[d][target]
}
