/*
 * @lc app=leetcode id=238 lang=golang
 *
 * [238] Product of Array Except Self
 */
func productExceptSelf(nums []int) []int {
	arr := make([]int, len(nums))
	arr[0] = 1
	for i:=1; i<len(nums); i++ {
		arr[i] = arr[i-1] * nums[i-1]
	}
	tmp := 1
	for i:=len(nums)-1; i>=0; i--{
		arr[i] = tmp * arr[i]
		tmp *= nums[i] 
	}
	return arr
}

