/*
 * @lc app=leetcode id=137 lang=golang
 *
 * [137] Single Number II
 */
// (4 ms) √ Your runtime beats 99.42 %
// 不带进位的3进制
func singleNumber(nums []int) int {
	x1, x2, mask := 0, 0, 0
	for _, i := range nums {
		x2 ^= (x1 & i)
		x1 ^= i
		mask = ^(x2 & x1)
		x2 &= mask
		x1 &= mask
	}
	return x1
}

// 模板：
//  一组数据，只有一个数出现p次，其他所有数字出现k次
// 找m，满足 2^m>=k,  本题中m=2，即需要两个变量 x1，x2. 如果2^m>k,则需要mask。本题目中0，1，2，0，1，2， 出现3的时候归零，0b11,mask = ~(x1&x2)
// 返回值看p，本题中p = 0b01，所以return x1

// for i := range nums{
// 	x2 ^= (x1&i)
// 	x1 ^= i
// 	mask = ~(x1&x2)
// 	x2 &= mask
// 	x1 &=mask
// }
