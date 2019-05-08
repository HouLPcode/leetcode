/*
 * @lc app=leetcode id=62 lang=golang
 *
 * [62] Unique Paths
 */
func uniquePaths(m int, n int) int {
	// f[i][j] 移动到块[i][j]的走法个数
	// f[i][j] = f[i-1][j] + f[i][j-1]
	// 结果 f[m][n]

	// 怎么创建二维[][]int ?????
	// 是否能直接同过缓存两个变量计算出结果？？？？？估计不行
	buf := make([]int,m*n,m*n)
	for i:=0;i<m;i++{
		buf[i] = 1
	}
	for j:=0;j<n;j++{
		buf[j*m] = 1
	}

	for i:=1;i<m;i++{
		for j:=1;j<n;j++{
			buf[i+j*m] = buf[i-1+j*m] + buf[i+(j-1)*m]
		}
	}
	return buf[m*n-1]
}

