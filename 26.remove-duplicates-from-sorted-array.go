/*
 * @lc app=leetcode id=26 lang=golang
 *
 * [26] Remove Duplicates from Sorted Array
 */
//  双指针
// p1：当前不重复的索引
// p2：遍历的索引。
func removeDuplicates(nums []int) int {
    if len(nums) == 0{
		return 0
	}
	p1 := 0
	for p2:=1; p2 < len(nums); p2++ {
		if nums[p1] != nums[p2]{
			nums[p1+1] = nums[p2]
			p1++
		}
	}
	return p1+1
}

