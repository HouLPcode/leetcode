/*
 * @lc app=leetcode id=790 lang=golang
 *
 * [790] Domino and Tromino Tiling
 */
// (0 ms) √ Your runtime beats 100 %
// 三种状态， 2*i的空间，
// dp[i][0] 最后边是竖的
// dp[i][1] 下边少出一块
// dp[i][2] 上边少出一块
// 转移方程
// dp[i][0] = dp[i-1][0] + dp[i-2][0] + dp[i-1][1] + dp[i-1][2]
// dp[i][1] = dp[i-2][0] + dp[i-1][2]
// dp[i][2] = dp[i-2][0] + dp[i-1][1]// 发现dp[i][1]和dp[i][2] 总是一样的，可以合并
// 最后结果 dp[N][0]
// 初始状态 dp[0][.] = 0  dp[1][.]=1

func numTilings(N int) int {
	if N == 1 {
		return 1
	}
	// 一定注意，此处不要是切片，会导致错误
	pri1 := [3]int{1, 0, 0} // 1
	pri2 := [3]int{2, 1, 1} // 2

	for i := 2; i < N; i++ {
		pri1, pri2[0], pri2[1], pri2[2] = pri2, (pri2[0]+pri1[0]+pri2[1]+pri2[2])%1000000007, (pri1[0]+pri2[2])%1000000007, (pri1[0]+pri2[1])%1000000007
	}
	return pri2[0]
}
