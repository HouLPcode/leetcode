/*
 * @lc app=leetcode id=50 lang=golang
 *
 * [50] Pow(x, n)
 */
func myPow(x float64, n int) float64 {
  if n < 0{
		return float64(1.0/myPow(x,-n))
	}
	// 非递归实现
	out := float64(1.0)
	for n != 0{
		// 注意一下3个计算步骤的顺序
		out *= out
		if n&1 == 1{
			out *= x // 奇数多乘一个x
		}
		n /= 2
	}
	return out
}

