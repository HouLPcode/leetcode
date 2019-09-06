/*
 * @lc app=leetcode.cn id=407 lang=golang
 *
 * [407] 接雨水 II
 *
 * https://leetcode-cn.com/problems/trapping-rain-water-ii/description/
 *
 * algorithms
 * Hard (31.26%)
 * Likes:    93
 * Dislikes: 0
 * Total Accepted:    909
 * Total Submissions: 2.9K
 * Testcase Example:  '[[1,4,3,1,3,2],[3,2,1,3,2,4],[2,3,3,2,3,1]]'
 *
 * 给定一个 m x n 的矩阵，其中的值均为正整数，代表二维高度图每个单元的高度，请计算图中形状最多能接多少体积的雨水。
 * 
 * 
 * 
 * 说明:
 * 
 * m 和 n 都是小于110的整数。每一个单位的高度都大于 0 且小于 20000。
 * 
 * 
 * 
 * 示例：
 * 
 * 给出如下 3x6 的高度图:
 * [
 * ⁠ [1,4,3,1,3,2],
 * ⁠ [3,2,1,3,2,4],
 * ⁠ [2,3,3,2,3,1]
 * ]
 * 
 * 返回 4。
 * 
 * 
 * 
 * 
 * 如上图所示，这是下雨前的高度图[[1,4,3,1,3,2],[3,2,1,3,2,4],[2,3,3,2,3,1]] 的状态。
 * 
 * 
 * 
 * 
 * 
 * 下雨后，雨水将会被存储在这些方块中。总的接雨水量是4。
 * 
 */
 //  × testcase: '[[12,13,1,12],[13,4,13,12],[13,8,10,12],[12,13,12,12],[13,13,13,13]]'
// err-------------------------------
 func trapRainWater(heightMap [][]int) int {
    rows := len(heightMap)
    if rows == 0 {
        return 0
    }
    cols := len(heightMap[0])
    if cols == 0 {
        return 0
    }
    
    left := make([][]int, rows)
    for i:=0; i<rows; i++{
        left[i] = make([]int, cols)
    }
    for i:=0; i<rows; i++ {
        left[i][0] = heightMap[i][0]
        for j:=1; j<cols; j++{
            if heightMap[i][j] > left[i][j-1]{
                 left[i][j] = heightMap[i][j]
            } else {
                left[i][j] = left[i][j-1]
            }
        }
    }
    
    right := make([][]int, rows)
    for i:=0; i<rows; i++{
        right[i] = make([]int, cols)
    }
    for i:=0; i<rows; i++ {
        right[i][cols-1] = heightMap[i][cols-1]
        for j:=cols-2; j>=0; j--{
            if heightMap[i][j] > right[i][j+1]{
                 right[i][j] = heightMap[i][j]
            } else {
                right[i][j] = right[i][j+1]
            }
        }
    }
    
    up := make([][]int, rows)
    for i:=0; i<rows; i++{
        up[i] = make([]int, cols)
    }
    for c:=0; c<cols;c++ {
        up[0][c] = heightMap[0][c]
        for r:=1; r<rows;r++{
            if heightMap[r][c] > up[r-1][c] {
                up[r][c] = heightMap[r][c]
            }else{
                up[r][c] = up[r-1][c]
            }
        }
    }
    
    down := make([][]int, rows)
    for i:=0; i<rows; i++{
        down[i] = make([]int, cols)
    }
    for c:=0; c<cols;c++ {
        down[rows-1][c] = heightMap[rows-1][c]
        for r:=rows-2; r>=0;r--{
            if heightMap[r][c] > down[r+1][c] {
                down[r][c] = heightMap[r][c]
            }else{
                down[r][c] = down[r+1][c]
            }
        }
    }
    sum := 0 
    for r:=1; r<rows-1;r++{
        for c:=1;c<cols-1;c++{
            if heightMap[r][c] < left[r][c-1] && 
            heightMap[r][c] < right[r][c+1] &&
            heightMap[r][c] < up[r-1][c] &&
            heightMap[r][c] < down[r+1][c]{
                sum += min4(left[r][c-1],right[r][c+1], up[r-1][c],down[r+1][c]) - heightMap[r][c]
            }
        }
    }
    return sum
}

func min4(a,b,c,d int) int{
    if a<=b&& a<= c && a<= d{
        return a
    }else if b<=a&& b<= c && b<= d{
        return b
    }else if c<=a&& c<= b && c<= d{
        return c
    }
    return d
}

