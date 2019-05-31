/*
 * @lc app=leetcode id=54 lang=golang
 *
 * [54] Spiral Matrix
 */
// [[1, 2, 3, 4],[5, 6, 7, 8],[9,10,11,12]] 
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
func print(matrix [][]int, sr,sc,er,ec int, out *[]int){
	// 上 
	for i:=sc; i<=ec; i++ {
		(*out) = append((*out), matrix[sr][i])
	}
	// 右
	for i:=sr+1; i<=er; i++ {
		(*out) = append((*out), matrix[i][ec])
	}
	// 下
	for i:=ec-1; i>=sc; i-- {
		(*out) = append((*out), matrix[er][i])
	}
	// 左
	for i:=er-1; i>sr; i-- { // 注意避免重复输出转角节点
		(*out) = append((*out), matrix[i][sc])
	}
}

