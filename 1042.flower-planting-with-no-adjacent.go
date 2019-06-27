
/*
 * @lc app=leetcode id=1042 lang=golang
 *
 * [1042] Flower Planting With No Adjacent
 */
//  172 ms 71.26 %
//  步骤1 paths -> graph
//  步骤2 循环 bfs 找到结果
func gardenNoAdj(N int, paths [][]int) []int {
	graph := make([][]int, N)
	for _, v := range paths {
		graph[v[0]-1] = append(graph[v[0]-1], v[1]-1)
		graph[v[1]-1] = append(graph[v[1]-1], v[0]-1)
	}
	used := make([]bool, N)
	rtn := make([]int, N)
	for i:=0; i<N; i++ {
		if used[i] == false {
			bfs11(graph, i, used, rtn)
		}
	}
	return rtn
}

// 此处bfs中的s指定访问的起始节点，图可能不是全联通的
func bfs11(graph [][]int, s int, used []bool, ans []int) {
	queue := []int{s}
	for len(queue) > 0 {
		l := len(queue)
		for i:=0; i<l; i++ {
			node := queue[0]
			queue = queue[1:]
			if used[node] {
				continue
			}
			// 给当前节点赋值
			vals := make([]bool, 4)
			for _, v := range graph[node] {
				if used[v] == false { //未访问的节点放入队列
					queue = append(queue, v)
				} else { // 已访问节点，记录花种类
					vals[ans[v]-1] = true
				}
			}
			for i:=0; i<4; i++ {
				if vals[i] == false {
					ans[node] = i+1 // 花的种类1-4
					break
				}
			}
			used[node] = true
		}
	}
}
