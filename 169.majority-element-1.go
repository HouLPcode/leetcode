/*
 * @lc app=leetcode id=169 lang=golang
 *
 * [169] Majority Element
 */
// (12 ms) √ Your runtime beats 99.32 %
// O(n)， 两个变量，一个统计数字，一个统计次数
func majorityElement(nums []int) int {
	// if len(nums) == 0 {
	rtn, cnt := nums[0], 1
	for i := 1; i < len(nums); i++ {
		if cnt == 0 { // 一定注意，先判断出现次数
			rtn = nums[i]
			cnt++
		} else if nums[i] == rtn {
			cnt++
		} else {
			cnt--
		}
	}
	return rtn
}
