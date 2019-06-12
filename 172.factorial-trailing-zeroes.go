/*
 * @lc app=leetcode id=172 lang=golang
 *
 * [172] Factorial Trailing Zeroes
 */
//  3388ms  5.45%
// log时间复杂度怎么计算解决？？？
// 技巧，统计1~n中5的个数，注意25=5*5是2个5
// 因为尾数0是2*5产生的，连续的数中5的个数肯定比2的个数少，所以统计5的个数即是尾数0的个数
func trailingZeroes(n int) int {
	cnt := 0 
	// 每5个数比较一次
	for i:=0; i<=n; i+=5 { // 一定注意此处不是逐步遍历到n
		cnt += have5(i)
	}
	return cnt
}

// a中含有5的个数， 25有2个5，125有3个5
func have5(a int) int{
	cnt := 0
	for a > 0 && a%5 == 0{
		cnt++
		a /= 5
	}
	return cnt
}

