/*
 * @lc app=leetcode id=119 lang=golang
 *
 * [119] Pascal's Triangle II
 */
func getRow(rowIndex int) []int {
	rtn := make([]int, rowIndex+1)
	tmp := make([]int, rowIndex+1)//存储上一层
	for row:=0; row<=rowIndex; row++ {
		rtn[0] = 1
		rtn[row] = 1
		for i:=1; i<row; i++ {
			rtn[i] = tmp[i-1] + tmp[i]
		}
		copy(tmp, rtn)
	}
	return rtn
}

