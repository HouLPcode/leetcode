/*
 * @lc app=leetcode id=53 lang=golang
 *
 * [53] Maximum Subarray
 */
func maxSubArray(nums []int) int {
	// f(i) 包含nums[i]的和的最大值
	// 连续， f(i) = max( f(i-1)+nums[i],nums[i])
	// 结果 max( f(0...i) )
	if nums == nil || len(nums) == 0{
		return 0
	}
	buf := make([]int,len(nums),len(nums))
	curMax := nums[0]
	for k,v := range nums{
		if k == 0{
			buf[k] = v
		}else if buf[k-1] > 0{
			buf[k] = buf[k-1]+v
		}else{
			buf[k] = v
		}
		if curMax < buf[k]{
			curMax = buf[k]
		}
	}
	return curMax
}

