/*
 * @lc app=leetcode id=997 lang=golang
 *
 * [997] Find the Town Judge
 */
//  104 ms 59.26 % 待优化
//  1\n[] ,0的存在导致的问题
// 构造有向图，表示方式临接矩阵，临接表
//  找入度N-1, 出度为0的下标
func findJudge(N int, trust [][]int) int {
	graph := make([][]int, N+1) // 1-N，0不用，方便计算
	// 构造图
	for _, v := range trust {
		graph[v[0]] = append(graph[v[0]], v[1]) // 0信任1
	}
	// 最多只有一个法官，所以先找出度0,然后找入度N-1
	outbuf := []int{}
	for k, v := range graph {
		if len(v) == 0 {
			outbuf = append(outbuf, k)
		}
	}
	for k, v := range outbuf {
		if k == 0 {
			continue
		}
		incnt := 0
		for _, val := range graph {
			for _, v1 := range val {
				if v1 == v {
					incnt++
				}
			}
		}
		if incnt == N-1 {
			return v
		}
	}
	return -1
}

