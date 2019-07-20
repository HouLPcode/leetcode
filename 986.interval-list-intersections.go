/*
 * @lc app=leetcode id=986 lang=golang
 *
 * [986] Interval List Intersections
 */
//  (40 ms) Your runtime beats 94.16 %
func intervalIntersection(A [][]int, B [][]int) [][]int {
	rtn := make([][]int, 0)
	lA, lB := len(A), len(B)
	i, j := 0, 0
	for i < lA && j < lB {
		if A[i][1] < B[j][0] || B[j][1] < A[i][0] {
		} else {
			rtn = append(rtn, []int{max(A[i][0], B[j][0]), min(A[i][1], B[j][1])})
		}
		if A[i][1] < B[j][1] {

			i++ // B 有剩余
		} else {
			j++
		}
	}

	return rtn
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
