/*
 * @lc app=leetcode id=441 lang=golang
 *
 * [441] Arranging Coins
 */
//  0 ms
func arrangeCoins(n int) int {
	// (1+r)*r/2 = n
	// 从 1-n 折半查找合适的r
	s, e := 1, n 
	for s <= e {
		mid := (e-s)/2 + s 
		if (1+mid)*mid / 2 <= n { //注意此处是小于等于
			s = mid + 1
		} else {
			e = mid -1
		}
	}
	return s-1 // 注意此处的-1
}

