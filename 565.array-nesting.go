/*
 * @lc app=leetcode id=565 lang=golang
 *
 * [565] Array Nesting
 */
// (8 ms) √ Your runtime beats 100 %
// 循环数组
func arrayNesting(nums []int) int {
	// 记录 调整为nums[i] = i 过程中间的步数
	rtn := 0
	for i := 0; i < len(nums); i++ {
		step := 0
		for nums[i] != i { // 没有重复元素
			nums[i], nums[nums[i]] = nums[nums[i]], nums[i]
			step++
		}
		if step > rtn {
			rtn = step
		}
	}
	return rtn + 1 // 需要的是元素的个数
}
