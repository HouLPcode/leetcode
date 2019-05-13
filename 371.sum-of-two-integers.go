/*
 * @lc app=leetcode id=371 lang=golang
 *
 * [371] Sum of Two Integers
 */
func getSum(a int, b int) int {
	// 不用加减符号计算两数之和
	// a^b 不带进位的加法
	// (a&b) << 1 进位处理
	rnt := a ^ b
	if (a&b)!= 0{
		return getSum(rnt,(a&b)<<1) //递归实现
	}
	return rnt
}

