/*
 * @lc app=leetcode id=372 lang=golang
 *
 * [372] Super Pow
 */
// 错误思路: 把b组成int，然后折半相乘 ****b可以无限大，导致超时。b可以无限大，无法转换成int数值
// 思路1：十进制方式 a^1234567 % k = (a^123456%k)^10 %k * (a^7 % k) % k
func superPow(a int, b []int) int {
	if len(b) == 1{
		return pow(a,b[0])
	}
	num1 := superPow(a,b[:len(b)-1])
	num1 = pow(num1, 10)
	return num1 * pow(a, b[len(b)-1]) % 1337
}

func pow(a,b int) int {
	rtn := 1
	for ;b>0;b-- {
		rtn = rtn%1337 * (a%1337)
	}
	return rtn%1337
}

