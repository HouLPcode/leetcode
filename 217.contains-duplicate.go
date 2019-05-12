/*
 * @lc app=leetcode id=217 lang=golang
 *
 * [217] Contains Duplicate
 */
func containsDuplicate(nums []int) bool {
	// 先排序，后查重
	sort.Ints(nums)
	for i:=0; i < len(nums)-1; i++{
		if nums[i] == nums[i+1]{
			return true
		}
	}
	return false
}

