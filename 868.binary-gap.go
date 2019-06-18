/*
 * @lc app=leetcode id=868 lang=golang
 *
 * [868] Binary Gap
 */
//  0 ms
// 注意0b11和0b10的区别
func binaryGap(N int) int {
	N /= (N & (-N)) // 去掉最右侧1后面的数
	if N == 1 {
		return 0
	}
	rtn := 0
	cnt := 0//连续0的个数
	for N != 0{
		if N & 1 == 0{
			cnt++
		} else {
			cnt = 0
		}
		if cnt > rtn {
			rtn = cnt
		}
		N >>= 1
	}
	return rtn+1
}

