/*
 * @lc app=leetcode id=200 lang=golang
 *
 * [200] Number of Islands
 */
 func numIslands(grid [][]byte) int { //传入的是byte有什么意思？？？
	// 法1 DFS
	// 法2 并查集
	if len(grid) == 0 {
		return 0
	} else if len(grid[0]) == 0 {
		return 0
	}
	rows, cols := len(grid), len(grid[0])
	roots := make([]int, rows*cols, rows*cols) //此处的len不能是0，否则后面的赋值只能用append
	for k := range roots {                     //初始化
		roots[k] = k
	}
	rank := make([]int, rows*cols, rows*cols)
	for i := 0; i < rows; i++ { // 只有一行或一列，2行，2列时候怎么处理
		for j := 0; j < cols; j++ {
			// 合并 每个节点的上下左右？？？ 移动数组
			if grid[i][j] == '1' {
				if i > 0 && grid[i-1][j] == '1' { //注意grid中是 '1'，此处应该传入i*j+j坐标
					union(&roots, &rank, i*cols+j, (i-1)*cols+j)
				}
				if i+1 < rows && grid[i+1][j] == '1' {
					union(&roots, &rank, i*cols+j, (i+1)*cols+j)
				}
				if j > 0 && grid[i][j-1] == '1' {
					union(&roots, &rank, i*cols+j, i*cols+j-1)
				}
				if j+1 < cols && grid[i][j+1] == '1' {
					union(&roots, &rank, i*cols+j, i*cols+j+1)
				}
				//union之后才能决定该点的parent
				//tmp[findParent(roots, i*cols+j)] = true //把该节点的parent加入，避免整图只有1个点的未统计
			}
		}
	}
	//统计个数，不能简单的统计tmp的长度
	tmp := make(map[int]bool)
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j]=='1'{
				tmp[findParent(roots, i*cols+j)] = true
			}
		}
	}
	return len(tmp)
}

func findParent(roots []int, x int) int {
	if roots[x] == x { //越界？？？
		return x
	} //TODO 优化：查询的过程中路径压缩
	return findParent(roots, roots[x]) //要return吗？？？
}

func union(roots *[]int, rank *[]int, x, y int) int {
	xp := findParent(*roots, x)
	yp := findParent(*roots, y)
	if xp != yp { //此处不能简单的赋值给左或右
		if (*rank)[xp] > (*rank)[yp] {
			(*roots)[yp] = xp
		} else if (*rank)[xp] < (*rank)[yp] {
			(*roots)[xp] = yp
		} else {
			(*roots)[yp] = xp
			(*rank)[xp]++
		}
	}
	//把parent返回 //注意是xp还是yp
	return findParent(*roots, xp) //可以不返回，在遍历的过程中统计parent的数量
}
