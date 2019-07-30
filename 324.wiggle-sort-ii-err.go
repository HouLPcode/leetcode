import "sort"

/*
 * @lc app=leetcode id=324 lang=golang
 *
 * [324] Wiggle Sort II
 */
// [4,5,5,6]
// 典型错误思路： 排序后折半成两个数组，然后隔一位插入一个数字
func wiggleSort(nums []int) {
	sort.Ints(nums)
	// 创建缓存
	nums1 := make([]int, (len(nums)+1)/2)
	nums2 := make([]int, len(nums)/2)
	copy(nums1, nums[:(len(nums)+1)/2])
	copy(nums2, nums[(len(nums)+1)/2:])

	index := 0
	for i := 0; i < len(nums1); i++ {
		nums[index] = nums1[i]
		index += 2
	}
	index = 1
	for i := 0; i < len(nums2); i++ {
		nums[index] = nums2[i]
		index += 2
	}
}
