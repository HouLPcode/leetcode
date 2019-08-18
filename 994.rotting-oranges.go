/*
 * @lc app=leetcode id=994 lang=golang
 *
 * [994] Rotting Oranges
 */
// 4 ms, faster than 87.50%
// BFS
// 不需要考虑多个起始节点，考虑成第一个节点有这么多子节点就行了，相当于少遍历一层
func orangesRotting(grid [][]int) int {
	badQueue := make([][2]int, 0) // 起始坏橘子个数
	rows, cols := len(grid), len(grid[0])
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] == 2 {
				badQueue = append(badQueue, [2]int{r, c})
			}
		}
	}

	time := 0 // 次数
	for len(badQueue) != 0 {
		queLen := len(badQueue)
		time++ // 遍历一层的时候次数++
		//if time == 20 { // 此处表示已经遍历完，最大的10*10最坏有18次，此处time多计算一个，所以time最大值为19，如果有20则表示有孤立的节点
		//return -1 // 没有重复节点不会进入循环中
		//}
		for i := 0; i < queLen; i++ {
			pix := badQueue[0]
			badQueue = badQueue[1:]
			for j := 0; j < 4; j++ {
				r := pix[0] + dirs[j][0]
				c := pix[1] + dirs[j][1]
				if r >= 0 && r < rows && c >= 0 && c < cols && grid[r][c] == 1 { // 只有好的才入队。避免重读计算
					badQueue = append(badQueue, [2]int{r, c})
					grid[r][c] = 2
				}
			}
		}
	}
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] == 1 {
				return -1
			}
		}
	}

	if time > 0 { // 注意全是0或1的情况
		time--
	}
	return time // 多遍历了一层，最后一层不会有任何入队操作
}

var dirs = [][]int{
	[]int{0, 1},
	[]int{0, -1},
	[]int{-1, 0},
	[]int{1, 0},
}
