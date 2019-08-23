/*
 * @lc app=leetcode id=1014 lang=golang
 *
 * [1014] Best Sightseeing Pair
 */
//  (52 ms) √ Your runtime beats 86.84 %
func maxScoreSightseeingPair(A []int) int {
	// res = A[i]+i + A[j]-j. (i < j)
	// 直接遍历一遍，过程种记录已知的 A[i]+i的最大值
	add := A[0] + 0
	rtn := 0
	for i := 1; i < len(A); i++ {
		if A[i]-i+add > rtn {
			rtn = A[i] - i + add
		}
		if A[i]+i > add {
			add = A[i] + i
		}
	}
	return rtn
}
