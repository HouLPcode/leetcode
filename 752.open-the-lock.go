/*
 * @lc app=leetcode id=752 lang=golang
 *
 * [752] Open the Lock
 */
//  48 ms 73.68 %  string -> int 优化nexts函数
//  132 ms 32.89 % 待优化
//  图的BFS搜索最短路径
// 4位数字，每个数字上下2个情况，一共对应8中情况
func openLock(deadends []string, target string) int {
	// map缓存，方便查找
	deadendsMap := make(map[int]bool, len(deadends))
	for _,v := range deadends {
		dig,_ := strconv.Atoi(v)
		deadendsMap[dig] = true
	}
	targetDig,_ := strconv.Atoi(target)
	return bfs(0, targetDig, deadendsMap) //通过deadendsMap当used使用？？？
}

// 广度优先， 非递归
func bfs(start, target int, used map[int]bool) int{
	step := -1 // 注意step初始值，否则可能多一个
	queue := []int{start}

	for len(queue) > 0 {
		l := len(queue)
		step++ // 一定注意是每层++
		for i:=0; i < l; i++ { // 访问一层节点,此处采用变量l而不是len，避免每次循环都重新获取queue长度的影响
			node := queue[0]
			queue = queue[1:]
			if used[node] == true {
				continue
			}
			used[node] = true
			if node == target {
				return step
			}
			for _, v := range nexts(node) {
				if used[v] == false {
					queue = append(queue, v)
				}
			}
		}
	}
	return -1
}

// 生成相邻的8个节点
func nexts(start int) []int {
	return []int{
		start%1000 + (start/1000+1)%10*1000,
		start%1000 + (start/1000-1+10)%10*1000,
		start/1000*1000 + start%100 + (start%1000/100+1)%10*100,
		start/1000*1000 + start%100 + (start%1000/100-1+10)%10*100,
		start/100*100 + start%10 + (start%100/10+1)%10*10,
		start/100*100 + start%10 + (start%100/10-1+10)%10*10,
		start/10*10 + (start%10+1)%10,
		start/10*10 + (start%10-1+10)%10,
	}
}
