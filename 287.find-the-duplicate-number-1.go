/*
 * @lc app=leetcode id=287 lang=golang
 *
 * [287] Find the Duplicate Number
 */
//  4 ms √ Your runtime beats 99.61 %
//  尝试调整数组元素，使 i 值在 i 索引位置上
// 注意题目中给出元素值是 1...n，不包括0
func findDuplicate(nums []int) int {
	for i := 0; i < len(nums); i++ {
		for i != nums[i] {
			if nums[i] == nums[nums[i]] { // 一定注意，此处判断是在for循环内，如果在for...i循环中，会出现死循环
				return nums[i] // 找到重复值
			}
			// 下面的交换并保证最终 i==nums[i],如果i是重复元素，则会一直在此处循环
			nums[i], nums[nums[i]] = nums[nums[i]], nums[i]
		}
	}
	return -1 // 不会执行此处，表示没有重复元素
}
