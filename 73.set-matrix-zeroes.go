/*
 * @lc app=leetcode id=73 lang=golang
 *
 * [73] Set Matrix Zeroes
 */
//  20ms 95.63%
// 用固定空间实现
// 用行首和列首标记，表示当前行列变成全0。由于左上到右下的过程中，不会重复访问标记后的行首列首，所以不会出现0后再次遍历的错误
func setZeroes(matrix [][]int)  {
	rows := len(matrix)
	if rows == 0 {
		return 
	}
	cols := len(matrix[0])
	colFlag := false
	// [0][0]是行首的标志，colFlag表示列首，因为如果第一行某元素是0，会将[0][0]赋值0，不能将其理解成列首也是0
	for i:=0; i<rows; i++ {
		if matrix[i][0] == 0{ 	// 处理列首元素
			colFlag = true
		}
		for j:=1; j<cols; j++ { // 从第1列开始 
			if matrix[i][j] == 0{
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}

	// 正确做法是从[1][1]往右下移动，根据行首列首标记决定是否重新赋值[i][j]
	for i:=1; i<rows; i++ {
		for j:=1; j<cols; j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			} 
		}
	}
	// 然后处理行首列首
	if matrix[0][0] == 0{
		for j:=0; j<cols; j++ {
			matrix[0][j] = 0
		}
	}
	if colFlag == true {
		for i:=0; i<rows; i++ {
			matrix[i][0] = 0
		}		
	}

	// 下面是典型错误，可能会导致所有元素均是0
	// for i:=0; i<rows; i++ {
	// 	if matrix[i][0] == 0{
	// 		for j:=0; j<cols; j++ {
	// 			matrix[i][j] = 0
	// 		}
	// 	}
	// }
	// for j:=0; j<cols; j++ {
	// 	if matrix[0][j] == 0{
	// 		for i:=0; i<rows; i++ {
	// 			matrix[i][j] = 0
	// 		}
	// 	}
	// }
}

