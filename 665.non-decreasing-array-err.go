/*
 * @lc app=leetcode id=665 lang=golang
 *
 * [665] Non-decreasing Array
 */
// [3,4,2,3]
// [2,3,3,2,4]
// [-1,4,2,3]
func checkPossibility(nums []int) bool {
	if len(nums) == 1 {
		return true
	}
	i := 0
	for i < len(nums)-1 {
		if nums[i] > nums[i+1] {
			break
		}
		i++
	}

	if i == len(nums)-1 { // 不用交换就是递增的
		return true
	}

	// i++ 典型错误：此处需要将i的值修改为i-1的值，不能i++，这样会忽略过去该值的检查
	// if i == 0 {
	// 	nums[i] = nums[i+1] // num[0]修改为num[1],其他的位置需要改为前面一个的值
	// } else {
	// 	nums[i] = nums[i-1] // 修改为前面的值
	// }
	if i == 0 {
		nums[i] = nums[i+1]
	} else {
		nums[i+1] = nums[i]
	}

	for i < len(nums)-1 {
		if nums[i] > nums[i+1] {
			break
		}
		i++
	}
	if i == len(nums)-1 { // 不用交换就是递增的
		return true
	}
	return false
}
