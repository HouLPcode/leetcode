/*
 * @lc app=leetcode id=198 lang=golang
 *
 * [198] House Robber
 */
func rob(nums []int) int {
	// DP
	//状态方程 f(i) 偷前i家的最大金额
	// f(i) = max(f(i-2)+nums[i], f(i-1))
	// 最终结果 f(n)
	if len(nums) == 0{
		return 0
	}
	// if len(nums) == 1{
	// 	return nums[0]
	// }
	pri0, pri1 := 0, nums[0]
	for i:=1; i<len(nums); i++ { //注意此处i从1开始
		pri0, pri1 = pri1, max(pri0+nums[i], pri1)
	}
	return pri1
}
func max(a, b int) int{
	if a > b {
		return a
	}
	return b
}

