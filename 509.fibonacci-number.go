/*
 * @lc app=leetcode id=509 lang=golang
 *
 * [509] Fibonacci Number
 */
func fib(N int) int {
	//不用处理N为负数
	pri1,pri2 := 0,1 //注意 顺序是0，1
	for ;N > 0; N--{ //注意 N>0
		pri1,pri2 = pri2,pri1+pri2
	}
	return pri1 //返回的是pri1
}

