/*
 * @lc app=leetcode id=41 lang=golang
 *
 * [41] First Missing Positive
 */
//  (0 ms)
// [1]
// [2,1]
// i+1放在索引i位置, 注意该题不会返回0，最小的正数是1
func firstMissingPositive(nums []int) int {
	for i := 0; i < len(nums); i++ {
		for i+1 != nums[i] {
			// 注意此处要判断索引为负，越界，重复值的情况
			if nums[i]-1 < 0 ||
				nums[i]-1 >= len(nums) ||
				nums[i] == nums[nums[i]-1] {
				break
			} else {
				nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
			}
		}
	}
	// [1]的处理
	i := 0
	for i < len(nums) && i == nums[i]-1 {
		i++
	}
	return i + 1
}
