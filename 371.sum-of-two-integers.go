/*
 * @lc app=leetcode id=371 lang=golang
 *
 * [371] Sum of Two Integers
 */
func getSum(a int, b int) int {
	// 不用加减符号计算两数之和
	// a^b 不带进位的加法
	// (a&b) << 1 进位处理
	return add2(a,b)
}

// 递归
func add1(a,b int)int{
	rnt := a ^ b
	if (a&b)!= 0{
		return add1(rnt,(a&b)<<1) //递归实现
	}
	return rnt
}

// 非递归
func add2(a,b int)int{
	cnt := a^b
	for a&b != 0{
		a,b = cnt, (a&b)<<1
		cnt = a ^ b
	}
	return cnt
}

