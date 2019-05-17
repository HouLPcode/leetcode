/*
 * @lc app=leetcode id=74 lang=golang
 *
 * [74] Search a 2D Matrix
 */
 // 二维数组摊平，按照一维sort二分查找
func searchMatrix(matrix [][]int, target int) bool {
    if len(matrix) == 0{
		return false
	}
	rows := len(matrix)
	cols := len(matrix[0])

	s,e := 0,rows*cols-1
	for s <= e{
		i,j := ito((e-s)/2+s, cols)
		if matrix[i][j] == target{
			return true
		} else if matrix[i][j] < target{
			s = (e-s)/2+s + 1
		}else{
			e = (e-s)/2+s -1
		}
	}
	return false
}

func toi(i,j,cols int)int{
	return i*cols+j
}

func ito(i,cols int) (int,int){
	return i / cols, i % cols
}

