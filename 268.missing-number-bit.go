/*
 * @lc app=leetcode id=268 lang=golang
 *
 * [268] Missing Number
 */
//  (16 ms) √ Your runtime beats 93.75 %
// 位运算
// n个元素对应  0,1,2,3,4,... n-1, n
// 数组下标对应 0,1,2,3,4,....n-1
// 所以 缺失的数字是 n ^ (index ^ nums[index]),不需要排序
func missingNumber(nums []int) int {
	rtn := len(nums) // 初始化为n
	for k, v := range nums {
		rtn ^= k
		rtn ^= v
	}
	return rtn
}
