/*
 * @lc app=leetcode id=986 lang=golang
 *
 * [986] Interval List Intersections
 */
//  100 ms 5.11 %
// 怎么确定合并的顺序？？？
func intervalIntersection(A [][]int, B [][]int) [][]int {
	rtn := make([][]int, 0)
	lA, lB := len(A), len(B)
	// 挨个比较，待优化
	for i := 0; i < lA; i++ {
		for j := 0; j < lB; j++ {
			if d, ok := merge(A[i], B[j]); ok {
				rtn = append(rtn, d)
			} else {
				// 合并不成功也不能 break，需要优化j起始比较的位置
			}
		}
	}
	// for i,j := 0,0; i<lA && j < lB; {
	//     if d,ok := merge(A[i], B[j]); ok{
	//         rtn = append(rtn, d)
	//         i++ //能够合并的时候访问A的下一个
	//     } else {
	//         i-- // 注意，不能合并的时候，需要和前一个i区间合并
	//         j++
	//     }
	// }
	// 剩余的不可能有交集
	return rtn
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
