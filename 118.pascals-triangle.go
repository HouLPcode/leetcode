/*
 * @lc app=leetcode id=118 lang=golang
 *
 * [118] Pascal's Triangle
 */
func generate(numRows int) [][]int {
	rtn := make([][]int, 0,numRows)
	for i:=0; i<numRows;i++{// 层数
		tmp := make([]int, i+1)// i层有i个元素
		tmp[0] = 1
		tmp[i] = 1
		for j:=1;j<i;j++{
			tmp[j] = rtn[i-1][j-1] + rtn[i-1][j]
		}
		rtn = append(rtn, tmp)
	}
	return rtn
}

