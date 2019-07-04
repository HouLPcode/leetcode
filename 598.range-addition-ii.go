/*
 * @lc app=leetcode id=598 lang=golang
 *
 * [598] Range Addition II
 */
//  4 ms  93.02%
// 每次都是从[0,0]开始，所以[0,0]的值最大，即找到和[0,0]值相同的元素，也就是每次ops中最小的值构成的范围元素个数
func maxCount(m int, n int, ops [][]int) int {
	minrow, mincol := m, n // 注意赋值的时候不用-1，因为在ops中是前闭后开的
	for _, v := range ops {
		minrow = min(minrow, v[0])
		mincol = min(mincol, v[1])
	}
	return minrow * mincol // 注意行列不用+1
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

