/*
 * @lc app=leetcode id=204 lang=golang
 *
 * [204] Count Primes
 */
 // 1不是素数，小于n
//  8ms beats 86.93%
func countPrimes(n int) int {
	// 从2开始，逐个遍历，将2，3，5，7，11...等质数的倍数全部标记为已访问
	count := 0
	buf := make([]bool, n)
	for i:=2; i < n; i++ {
		if buf[i] == false { //从2开始，第一次没有访问过的，一定是质数
			count++
			for j:=1; j*i<n; j++{//把i的整数倍标记为已访问，肯定不是素数
				buf[j*i] = true
			}
		}
	}
	return count
}

