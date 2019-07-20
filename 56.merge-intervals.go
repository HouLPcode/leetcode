import "sort"

/*
 * @lc app=leetcode id=56 lang=golang
 *
 * [56] Merge Intervals
 */
//  8 ms, faster than 97.10%
// 按照开始节点排序，之后两辆合并
// 使用额外的返回空间，而不是直接合并intervals
func merge(intervals [][]int) [][]int {
	rtn := make([][]int, 0, len(intervals))
	if len(intervals) == 0 {
		return rtn
	}
	// [][]int怎么排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	rtn = append(rtn, intervals[0]) // 先把第0个放进去
	// for i := 1; i < len(intervals); i++ {
	// 	if intervals[i][0] > rtn[len(rtn)-1][1] { // 一定注意，是intervals和rtn的比较，不是intervals之间的比较
	// 		// 不重合，直接加入
	// 		rtn = append(rtn, intervals[i])
	// 	} else {
	// 		// 合并，只需要把尾巴节点更新为最大值即可
	// 		if rtn[len(rtn)-1][1] < intervals[i][1] {
	// 			rtn[len(rtn)-1][1] = intervals[i][1]
	// 		}
	// 	}
	// }
	curEnd := rtn[0][1] // 优化，记录一直的最后节点，避免每次获取该值
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] > curEnd {
			// 不重合，直接加入
			rtn = append(rtn, intervals[i])
			curEnd = intervals[i][1]
		} else {
			// 合并，只需要把尾巴节点更新为最大值即可
			if curEnd < intervals[i][1] {
				rtn[len(rtn)-1][1] = intervals[i][1]
				curEnd = intervals[i][1]
			}
		}
	}

	return rtn
}
