/*
 * @lc app=leetcode id=48 lang=golang
 *
 * [48] Rotate Image
 */
 // 按圈旋转
func rotate(matrix [][]int)  {
	n := len(matrix)
	for i:=0; i<n/2 ; i++ { // 圈数，不回出现1个元素是1圈的情况
		rotate1(&matrix, i,i,n-1-i,n-1-i)
	}
	// rotate1(&matrix, 0,0,2,2)
}

// 某圈旋转,起始和终止的坐标
func rotate1(matrix *[][]int, sr,sc,er,ec int){
	// 旋转的规律	
	// 20 -> 00
	// 22 -> 20
	// 02 -> 22
	// 00 -> 02
	for i:=0; i< er-sr; i++{ //一组一组来，一共这么多组
		tmp := (*matrix)[sr][sc+i]
		(*matrix)[sr][sc+i] = (*matrix)[er-i][sc]
		(*matrix)[er-i][sc] = (*matrix)[er][ec-i]
		(*matrix)[er][ec-i] = (*matrix)[sr+i][ec]
		(*matrix)[sr+i][ec] = tmp
	}
}

