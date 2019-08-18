/*
 * @lc app=leetcode id=1091 lang=golang
 *
 * [1091] Shortest Path in Binary Matrix
 */
// (36 ms) √ Your runtime beats 100 %
// BFS
func shortestPathBinaryMatrix(grid [][]int) int {
	if grid[0][0] == 1 || grid[len(grid)-1][len(grid)-1] == 1 {
		return -1
	}
	depth := 0
	grid[0][0] = 1
	queue := [][2]int{[2]int{0, 0}}
	x, y := 0, 0
	for len(queue) != 0 {
		l := len(queue)
		depth++
		for i := 0; i < l; i++ {
			node := queue[0]
			queue = queue[1:]

			if node[0] == len(grid)-1 && node[1] == len(grid)-1 { //右下角元素，直接返回
				return depth
			}
			for i := 0; i < 8; i++ {
				x = node[0] + dirs[i][0]
				y = node[1] + dirs[i][1]
				if x >= 0 && x < len(grid) && y >= 0 && y < len(grid) && grid[x][y] == 0 {
					grid[x][y] = 1 // 入队的时候就标记该点，避免重复入队导致的超时或空间越界
					queue = append(queue, [2]int{x, y})
				}
			}
		}
	}
	return -1
}

var dirs = [][]int{
	[]int{-1, -1},
	[]int{-1, 0},
	[]int{-1, 1},
	[]int{0, -1},
	[]int{0, 1},
	[]int{1, -1},
	[]int{1, 0},
	[]int{1, 1},
}
