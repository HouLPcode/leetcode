/*
 * @lc app=leetcode id=9 lang=golang
 *
 * [9] Palindrome Number
 */
func isPalindrome(x int) bool {
	// 题目要求不能 int->string
	if x < 0{
		return false
	}
	// 尝试1：不断%,/ 组成新的数字，比较数值是否一致
	tmp := x
	cnt := tmp%10// 0也可以处理
	tmp /= 10	
	for tmp != 0{
		cnt = cnt*10 + tmp%10
		tmp /= 10
	}
	return x == cnt
}

