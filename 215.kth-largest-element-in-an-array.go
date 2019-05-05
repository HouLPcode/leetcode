/*
 * @lc app=leetcode id=215 lang=golang
 *
 * [215] Kth Largest Element in an Array
 */

func findKthLargest(nums []int, k int) int {
	// 思路1: 先排序，后输出.golang有两种排序方式sort.Slice和sort.Sort
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})
	return nums[k-1]
}

