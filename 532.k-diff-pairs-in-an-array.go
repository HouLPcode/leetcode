import "sort"

/*
 * @lc app=leetcode id=532 lang=golang
 *
 * [532] K-diff Pairs in an Array
 */
//  24 ms, faster than 61.29%
//  [1, 3, 1, 5, 4]\n0
func findPairs(nums []int, k int) int {
	// 先排序，相邻数字的差值最小，数字有正负
	if len(nums) < 2 {
		return 0
	}
	sort.Ints(nums)
	p0, p1 := 0, 1
	cnt := 0
	for p1 < len(nums) {
		if p0 == p1 {
			p1++ // 计算的时候两个指针不能指向同一个数
		} else if nums[p1]-nums[p0] == k {
			// 重复的解不计数，所以两个指针都右移到不同的值为止
			cnt++
			p0++ // 注意此处，移动有再比较，避免p0，p1原地不动
			p1++
			for ; p1 < len(nums) && nums[p1-1] == nums[p1]; p1++ {
			}
			for ; p0 < len(nums) && nums[p0-1] == nums[p0]; p0++ {
			}
		} else if nums[p1]-nums[p0] < k {
			p1++
		} else {
			p0++
		}
	}
	return cnt
}
