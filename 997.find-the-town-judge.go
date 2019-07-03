/*
 * @lc app=leetcode id=997 lang=golang
 *
 * [997] Find the Town Judge
 */
//  96 ms  100 %
//  1\n[] ,0的存在导致的问题
// 遍历的过程中统计每个节点的入度出度，不构造图
//  找入度N-1, 出度为0的下标
func findJudge(N int, trust [][]int) int {
	graph := make([][]int, N+1) // 1-N，0不用，方便计算
	for i := 1; i < N+1; i++ {
		graph[i] = make([]int, 2)
	}

	for _, v := range trust {
		graph[v[0]][0]++ //out
		graph[v[1]][1]++ // in
	}
	for i := 1; i < N+1; i++ {
		if graph[i][0] == 0 && graph[i][1] == N-1 {
			return i
		}
	}
	return -1
}

