/*
 * @lc app=leetcode id=162 lang=golang
 *
 * [162] Find Peak Element
 */
func findPeakElement(nums []int) int {
	// TODO log n 的做法是？？？？
	// 下面是O(n)做法，一次遍历，寻找到符合条件的值
	// 第一个？？？
	
	for i := 1; i < len(nums) - 1; i++{
		if nums[i] > nums[i-1] && nums[i] > nums[i+1]{
			return i
		}
	}

	if len(nums) > 1 && nums[len(nums)-1] > nums[len(nums)-2]{
		return len(nums) - 1
	}
	return 0 //前面的多不是，只能是第0个，因为肯定有结果
}

