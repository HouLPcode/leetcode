/*
 * @lc app=leetcode id=53 lang=golang
 *
 * [53] Maximum Subarray
 */
func maxSubArray(nums []int) int {
	// f(i) 包含nums[i]的和的最大值
	// 连续， f(i) = max( f(i-1)+nums[i],nums[i]) 
	// 结果 max( f(0...i) )
	// 转移方程只和前一个数有关，只缓存前一个计算结果即可
	if nums == nil || len(nums) == 0{
		return 0
	}
	curMax, sum := nums[0], nums[0]
	for _,v := range nums[1:]{
		sum = max(sum+v,v)
		curMax = max(curMax,sum)
	}
	return curMax
}
func max(a,b int)int{
	if a > b{
		return a
	}
	return b
}

