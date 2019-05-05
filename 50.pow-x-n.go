/*
 * @lc app=leetcode id=50 lang=golang
 *
 * [50] Pow(x, n)
 */
func myPow(x float64, n int) float64 {
  if n < 0{
		return float64(1.0/myPow(x,-n))
	}
	//忘记了递归结束条件
	if n==0{ // n==1 时多递归一次即可
		return 1
	}

	if n&1 == 0{//偶数
		// return float64(myPow(x,n/2) * myPow(x,n/2))// 这种方式有大量的重复运算，导致超时
		return float64(myPow(x*x,n/2))
	}
	// return float64(myPow(x,n/2) * myPow(x,n/2) * x)
	return float64(myPow(x*x,n/2) * x)
}

