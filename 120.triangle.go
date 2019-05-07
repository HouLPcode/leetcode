/*
 * @lc app=leetcode id=120 lang=golang
 *
 * [120] Triangle
 */
func minimumTotal(triangle [][]int) int {
	//状态方程: f(i,j) 从最底到(i,j)点的最小路径和
	//状态转移方程: f(i,j) = triangle[i][j] + min( f(i+1,j), f(i+1,j+1) )
	//所求值为f(0,0)
	//已知道 f(最低,0...j) = triangle[最低][0...j] 

	//???怎么range二维slice
	rows,columns := len(triangle),len(triangle[len(triangle)-1])
	rnt := help(0,0,rows,columns,triangle)
	return rnt
}

// 递归实现
func help(i,j,rows,cols int, nums [][]int) int{
	if i == rows - 1{
		return nums[i][j]
	}
	n1 := help(i+1,j,rows,cols,nums)
	n2 := help(i+1,j+1,rows,cols,nums)
	if n1 < n2{
		return nums[i][j] + n1
	}
	return nums[i][j] + n2 
}

