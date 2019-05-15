/*
 * @lc app=leetcode id=547 lang=golang
 *
 * [547] Friend Circles
 */
 // 拷贝自 leetcode 200 
 // 不能完全看成是岛屿数的题目，岛屿的关系是相邻，朋友的关系是同行同列中是1，可以相邻，也可以不相邻
func findCircleNum(M [][]int) int {
	if len(M) == 0 {
		return 0
	} else if len(M[0]) == 0 {
		return 0
	}
	rows, cols := len(M), len(M[0])
	roots := make([]int, rows*cols, rows*cols) //此处的len不能是0，否则后面的赋值只能用append
	for k := range roots {                     //初始化
		roots[k] = k
	}
	rank := make([]int, rows*cols, rows*cols)
	for i := 0; i < rows; i++ { // 只有一行或一列，2行，2列时候怎么处理
		for j := 0; j < cols; j++ {
			//把整行整列全整一遍，先不考虑去重
			// == 1
			if M[i][j] == 1{
				for k:=0;k<cols;k++{
					if M[i][k] == 1{
						union(&roots, &rank, i*cols+j, (i)*cols+k)//整行
					}
				}
				for k:=0;k<cols;k++{
					if M[k][j] == 1{
						union(&roots, &rank, i*cols+j, (k)*cols+j)//整列
					}
				}
			}
		}
	}
	//统计个数，不能简单的统计tmp的长度
	tmp := make(map[int]bool)
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if M[i][j]==1{
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

