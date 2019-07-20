/*
 * @lc app=leetcode id=50 lang=golang
 *
 * [50] Pow(x, n)
 */
// 正常情况下，判断底数，指数分别为负数，正数，0的情况
func myPow(x float64, n int) float64 {
	if n < 0 {
		return float64(1.0 / myPow(x, -n))
	}
	// 非递归实现
	out := float64(1.0)
	//注意此处的实现方式，x每次平方，n每次减半，为奇数时获取输出值
	for n != 0 {
		if n&1 == 1 {
			out *= x // 奇数多乘一个x
		}
		x *= x
		n >>= 1
	}
	return out
}
