/*
 * @lc app=leetcode id=120 lang=golang
 *
 * [120] Triangle
 */
// 技巧：该题状态方程定义的时候是从低往顶找的
// 如果从上往下找，需要计算最后一行中的最小值，从下往上算，直接返回顶部节点值即可
// 其实，不用考虑成DP问题，人脑计算的时候也是从下往上的过程中选择小的数计算到顶部
func minimumTotal(triangle [][]int) int {
	//状态方程: f(i,j) 从最底到(i,j)点的最小路径和
	//状态转移方程: f(i,j) = triangle[i][j] + min( f(i+1,j), f(i+1,j+1) )
	//所求值为f(0,0)
	//已知道 f(最低,0...j) = triangle[最低][0...j] 

	//怎么range二维slice
	rows,columns := len(triangle),len(triangle[len(triangle)-1])
	mems := make(map[string]int)
	rnt := help(0,0,rows,columns,triangle,mems)
	return rnt
}

//TODO 时间空间优化

// 递归实现，超时
func help(i,j,rows,cols int, nums [][]int, mems map[string]int) int{
	if v,ok := mems[fmt.Sprint(i,"+",j)]; ok{
		return v
	}
	if i == rows - 1{
		mems[fmt.Sprint(i,"+",j)] = nums[i][j]
		return nums[i][j]
	}
	n1 := help(i+1,j,rows,cols,nums,mems)
	n2 := help(i+1,j+1,rows,cols,nums,mems)
	if n1 < n2{
		mems[fmt.Sprint(i,"+",j)] = nums[i][j] + n1
		return nums[i][j] + n1
	}
	mems[fmt.Sprint(i,"+",j)] = nums[i][j] + n2
	return nums[i][j] + n2 
}

