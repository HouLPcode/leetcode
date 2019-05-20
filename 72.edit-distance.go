/*
 * @lc app=leetcode id=72 lang=golang
 *
 * [72] Edit Distance
 */
func minDistance(word1 string, word2 string) int {
	// DP
	// dp[i][j] word1前i字母变成word2前j字母的最小操作数
	// if w1[i] == w2[i] dp[i][j] = dp[i-1][j-1]
	// else   dp[i][j] = 1 + min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1])
	//                            insert         Delete      Replace
	// 最终结果: dp[m][n]

	// 初始条件 dp[i][0] = i, dp[0][j] = j
	m, n := len(word1), len(word2)
	dp := make([][]int, m+1) //注意是m+1
	for i:=0; i<m+1; i++{
		dp[i] = make([]int, n+1)
	}

	for i:=1; i<m+1; i++{
		dp[i][0] = i
	}
	for j:=1; j<n+1; j++{
		dp[0][j] = j
	}
	
	for i:=1; i<m+1; i++ {
		for j:=1; j<n+1; j++{
			if word1[i-1] == word2[j-1]{// 此处需要-1
				dp[i][j] = dp[i-1][j-1] 
			}else{
				dp[i][j] = 1 + min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1])
			}
		}
	}
	return dp[m][n]
}

func min(a, b, c int) int{
	if a<=b && a<=c{
		return a
	}else if b<=a && b<=c{
		return b
	}else{
		return c
	}
}

