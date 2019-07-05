import "math"

/*
 * @lc app=leetcode id=492 lang=golang
 *
 * [492] Construct the Rectangle
 */
//  28 ms  20.45 % 通过减少比较和计算量进行优化
// 开方，++ --
// func constructRectangle(area int) []int {
// 	rtn := []int{ // [L,W]
// 		int(math.Sqrt(float64(area))),
// 		int(math.Sqrt(float64(area))),
// 	}
// 	for {
// 		if rtn[0]*rtn[1] > area {
// 			rtn[1]--
// 		} else if rtn[0]*rtn[1] < area {
// 			rtn[0]++
// 		} else {
// 			break
// 		}
// 	}
// 	return rtn
// }

// 0 ms
func constructRectangle(area int) []int {
	w := int(math.Sqrt(float64(area)))
	for area%w != 0 {
		w--
	}
	return []int{area / w, w}
}

