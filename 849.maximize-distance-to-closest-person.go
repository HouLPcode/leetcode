/*
 * @lc app=leetcode id=849 lang=golang
 *
 * [849] Maximize Distance to Closest Person
 */
//  (12 ms) √ Your runtime beats 91.53 %
// [1,0,0,1,0,1,0,1]
func maxDistToClosest(seats []int) int {
	left, right := 0, 0 // 左右到第一个1为止，0的个数
	for i := 0; i < len(seats) && seats[i] != 1; i++ {
		left++
	}
	for i := len(seats) - 1; i >= 0 && seats[i] != 1; i-- {
		right++
	}

	curr := 0 // 统计中间1的最大间隔
	for s, e := left, left+1; e < len(seats)-right; e++ {
		if seats[e] == 1 {
			if e-s-1 > curr {
				curr = e - s - 1 // 中间0的个数
			}
			s = e
		}
	}

	rtn := 0
	// 返回3个数的最大情况
	if left > right {
		rtn = left
	} else {
		rtn = right
	}
	if rtn < (curr+1)/2 {
		rtn = (curr + 1) / 2
	}
	return rtn
}
