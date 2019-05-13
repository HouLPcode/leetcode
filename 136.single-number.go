/*
 * @lc app=leetcode id=136 lang=golang
 *
 * [136] Single Number
 */
func singleNumber(nums []int) int {
    //O(n)  O(1)
	// 技巧：所有数字异或，最终结果即为只出现一次的数字
	cnt := 0
	for _,v := range nums{
		cnt ^= v
	}
	return cnt
}
