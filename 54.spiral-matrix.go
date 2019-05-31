/*
 * @lc app=leetcode id=54 lang=golang
 *
 * [54] Spiral Matrix
 */
// 测试 正方形，长方形横，竖 三种情况 
func spiralOrder(matrix [][]int) []int {
	rows := len(matrix)
	if rows == 0{
		return []int{}
	}
	cols := len(matrix[0])
	out := make([]int, 0, rows*cols)
	// 按圈打印
	for i:=0; i<=rows-1-i && i<=cols-1-i; i++{
		print(matrix, i,i,rows-1-i,cols-1-i, &out)
	}
	return out
}

// 根据左上和右下顺时针打印一圈
// 优先闭区间
func print(matrix [][]int, sr,sc,er,ec int, out *[]int){
	for i:=sc; i<=ec; i++ {// 上，左闭右闭，注意避免重复输出转角节点
		(*out) = append((*out), matrix[sr][i])
	}
	for i:=sr+1; i<=er; i++ {// 右，上开下开
		(*out) = append((*out), matrix[i][ec])
	}
	if sr < er && sc < ec { //一定注意此处的判断，避免一行或一列的影响
		for i:=ec-1; i>=sc; i-- {// 下
			(*out) = append((*out), matrix[er][i])
		}
		for i:=er-1; i>sr; i-- { // 左 
			(*out) = append((*out), matrix[i][sc])
		}
	}
}

