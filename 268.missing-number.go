/*
 * @lc app=leetcode id=268 lang=golang
 *
 * [268] Missing Number
 */
func missingNumber(nums []int) int {
	// 数组长度l，和为(1+l)*l/2,减速sum
	sum,l,v := 0,0,0
	for l,v = range nums{
		sum += v
	}
	return (l+2)*(l+1)/2 - sum
}

