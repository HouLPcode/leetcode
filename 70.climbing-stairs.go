/*
 * @lc app=leetcode id=70 lang=golang
 *
 * [70] Climbing Stairs
 */
func climbStairs(n int) int {
	//状态定义: f(n) 跳到第n阶的走法
	//转移方程: f(n) = f(n-1) + f(n-2)
	// 题目指定 n为正
	
	// if n == 1 {
	// 	return 1
	// }
	// if n == 2{
	// 	return 2
	// }
	// 递归方式，超时
	// return climbStairs(n-1) + climbStairs(n-2)
	
	pri1, pri2 := 1, 1 // 注意，起始都为1，方便对输入1，2的处理
	for i := 2; i <= n; i++{ //注意，从2开始，小于等于号
		pri1, pri2 = pri2, pri1 + pri2 //多值同时赋值
	}
	return pri2
}

