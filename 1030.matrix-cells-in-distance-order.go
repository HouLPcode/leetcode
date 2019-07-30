/*
 * @lc app=leetcode id=1030 lang=golang
 *
 * [1030] Matrix Cells in Distance Order
 */
//  (544 ms) √ Your runtime beats 79.59 %
func allCellsDistOrder(R int, C int, r0 int, c0 int) [][]int {
	// 最大的距离是 R+C
	rtn := make([][]int, 0, R*C)
	rtn = append(rtn, []int{r0, c0})
	for i := 1; i < R+C-1; i++ {
		// 此处对0单独处理，避免+0 -0 导致的重复解
		if c0+i >= 0 && c0+i < C {
			rtn = append(rtn, []int{r0, c0 + i})
		}
		if c0-i >= 0 && c0-i < C {
			rtn = append(rtn, []int{r0, c0 - i})
		}
		if r0+i >= 0 && r0+i < R {
			rtn = append(rtn, []int{r0 + i, c0})
		}
		if r0-i >= 0 && r0-i < R {
			rtn = append(rtn, []int{r0 - i, c0})
		}

		for j := 1; j < i; j++ { // 注意此处没有==，避免重复
			if r0+j >= 0 && r0+j < R && c0+i-j >= 0 && c0+i-j < C {
				rtn = append(rtn, []int{r0 + j, c0 + i - j})
			}
			if r0+j >= 0 && r0+j < R && c0-i+j >= 0 && c0-i+j < C {
				rtn = append(rtn, []int{r0 + j, c0 - i + j})
			}
			if r0-j >= 0 && r0-j < R && c0+i-j >= 0 && c0+i-j < C {
				rtn = append(rtn, []int{r0 - j, c0 + i - j})
			}
			if r0-j >= 0 && r0-j < R && c0-i+j >= 0 && c0-i+j < C {
				rtn = append(rtn, []int{r0 - j, c0 - i + j})
			}
		}
	}
	return rtn
}
