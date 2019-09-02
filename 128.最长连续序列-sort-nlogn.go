import "sort"

/*
 * @lc app=leetcode.cn id=128 lang=golang
 *
 * [128] 最长连续序列
 */
//  8 ms, faster than 46.27%
// 法1：排序后查找，O(nlogn),不满足题目要求
func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	sort.Ints(nums)

	curLen := 1
	maxLen := 1
	for i := 1; i < len(nums); i++ { // 从1开始计算
		if nums[i] == nums[i-1] {
			// 重复元素，计数不更新
		} else if nums[i] == nums[i-1]+1 {
			curLen++
		} else {
			if curLen > maxLen {
				maxLen = curLen
			}
			curLen = 1
		}
	}
	if curLen > maxLen { // 统计最后一段的长度
		maxLen = curLen
	}
	return maxLen
}
