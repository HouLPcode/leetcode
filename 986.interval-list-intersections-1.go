/*
 * @lc app=leetcode id=986 lang=golang
 *
 * [986] Interval List Intersections
 */
// 44 ms, faster than 57.66%
// 怎么确定合并的顺序？？？
func intervalIntersection(A [][]int, B [][]int) [][]int {
	rtn := make([][]int, 0)
	lA, lB := len(A), len(B)
	j := 0
	i := 0
	for i < lA && j < lB {
		if d, ok := merge(A[i], B[j]); ok {
			rtn = append(rtn, d)
		}
		if A[i][1] < B[j][1] {
			i++ // B有剩余
		} else {
			j++ // A有剩余或都没剩余
		}
	}

	return rtn // 剩余的不可能有交集
}

func merge(a, b []int) ([]int, bool) {
	if a[1] < b[0] || b[1] < a[0] {
		return []int{}, false
	} else {
		return []int{max(a[0], b[0]), min(a[1], b[1])}, true
	}
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