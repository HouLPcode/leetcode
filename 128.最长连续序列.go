/*
 * @lc app=leetcode.cn id=128 lang=golang
 *
 * [128] 最长连续序列
 */
//  4 ms, faster than 96.52%
// 法1：排序后查找，O(nlogn),不满足题目要求
// 法2： map缓存。下面是实现方式
func longestConsecutive(nums []int) int {
	if len(nums) == 0 || len(nums) == 1 {
		return len(nums)
	}

	buf := make(map[int]bool, len(nums))
	for i := 0; i < len(nums); i++ {
		buf[nums[i]] = true
	}

	curLen := 0
	maxLen := 1
	for _, v := range nums {
		if buf[v-1] { // 一定注意，此处是对v-1进行判断，如果不存在，表示v是一个序列的开始
			continue
		}
		curLen = 0
		for buf[v] {
			v++
			curLen++
		}
		if curLen > maxLen {
			maxLen = curLen
		}
	}
	if curLen > maxLen {
		maxLen = curLen
	}
	return maxLen
}
