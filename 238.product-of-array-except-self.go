/*
 * @lc app=leetcode id=238 lang=golang
 *
 * [238] Product of Array Except Self
 */
func productExceptSelf(nums []int) []int {
	// 题目中要求num长度大于1，且输出空间不算是额外开辟的空间
	arr := make([]int, len(nums)) // 存储前i项乘积，不包含第i项
	arr[0] = 1
	for i:=1; i<len(nums); i++ {
		arr[i] = arr[i-1] * nums[i-1]
	}
	tmp := 1 // 从右往左的乘积
	for i:=len(nums)-1; i>=0; i--{
		arr[i] = tmp * arr[i]
		tmp *= nums[i] 
	}
	return arr
}

