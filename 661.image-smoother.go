/*
 * @lc app=leetcode id=661 lang=golang
 *
 * [661] Image Smoother
 */
/*
 * @lc app=leetcode id=661 lang=golang
 *
 * [661] Image Smoother
 */
//  80 ms  25.93 %
// 重新创建返回结果，不使用M
func imageSmoother(M [][]int) [][]int {
	rows := len(M)
	if rows == 0 {
		return [][]int{}
	}
	cols := len(M[0])
	rtn := make([][]int, 0, rows)
	for r := 0; r < rows; r++ {
		rtn = append(rtn, make([]int, 0, cols)
		for c := 0; c < cols; c++ {
			avg := smooth(M, r, c, rows, cols)
			rtn[r] = append(rtn[r], avg)
		}
	}
	return rtn
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

func smooth(M [][]int, r, c, rows, cols int) int {
	sum := M[r][c]
	cnt := 1
	for i := 0; i < 8; i++ {
		if r+dirs[i][0] >= 0 && r+dirs[i][0] < rows &&
			c+dirs[i][1] >= 0 && c+dirs[i][1] < cols {
			sum += M[r+dirs[i][0]][c+dirs[i][1]]
			cnt++
		}
	}
	return sum / cnt
}
