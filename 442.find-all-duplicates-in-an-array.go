/*
 * @lc app=leetcode id=442 lang=golang
 *
 * [442] Find All Duplicates in an Array
 */
// (388 ms) √ Your runtime beats 70.54 %
func findDuplicates(nums []int) []int {
	// i位置存储i
	rtn := []int{}
	for i := 1; i <= len(nums); i++ {
		for nums[i-1] != i && nums[i-1] != 0 {
			if nums[i-1] == nums[nums[i-1]-1] {
				rtn = append(rtn, nums[i-1])
				nums[i-1] = 0 // 重复的值赋值为0
				break
			}
			nums[i-1], nums[nums[i-1]-1] = nums[nums[i-1]-1], nums[i-1]
		}
	}
	return rtn
}
