/*
 * @lc app=leetcode id=213 lang=golang
 *
 * [213] House Robber II
 */
 // leetcode 198 数组换成环
func rob(nums []int) int {
	// 第1家偷，最后一家肯定不偷，返回f(n-1)
	// 第1家不偷，最后一家可以偷，f(n)
	// 返回f(n-1)和f(n)的最大值

	if len(nums) == 0{
		return 0
	}
	if len(nums) == 1{
		return nums[0]
	}

	// 偷第1家
	pri0, pri1 := 0, nums[0]
	for i:=1; i<len(nums)-1; i++ { //注意此处i从1开始，到n-1肯定不能偷，到n-2位置
		pri0, pri1 = pri1, max(pri0+nums[i], pri1)
	}
	// 不偷第1家
	pri0, pri2 := 0, nums[1]//只有1家此处越界
	for i:=2; i<len(nums); i++ { //注意此处i从1开始
		pri0, pri2 = pri2, max(pri0+nums[i], pri2)
	}
	return max(pri1, pri2)

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

