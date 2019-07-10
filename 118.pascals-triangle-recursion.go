/*
 * @lc app=leetcode id=118 lang=golang
 *
 * [118] Pascal's Triangle
 */
//  0 ms
 // 递归
 func generate(numRows int) [][]int {
    rtn := make([][]int, 0, numRows)
    for r:=0; r < numRows; r++ {
        rtn = append(rtn, make([]int, r+1))
        rtn[r][0] = 1
        rtn[r][r] = 1
    }
    for c := 0; c < numRows; c++ {
        fill(&rtn, numRows-1, c)
    }
    return rtn
}

func fill(buf *[][]int, r, c int) int{
    if (*buf)[r][c] > 0 { // 剪枝，避免重复计算
        return (*buf)[r][c]
    }
    (*buf)[r][c] = fill(buf, r-1, c-1) + fill(buf, r-1, c)
    return (*buf)[r][c]
}

