/*
 * @lc app=leetcode id=1011 lang=golang
 *
 * [1011] Capacity To Ship Packages Within D Days
 */
//  (36 ms) √ Your runtime beats 76.74 %
func shipWithinDays(weights []int, D int) int {
	// 统计出最大最小的范围，然后二分法查找最小值是否满足条件
	min := 501
	sum := 0
	for _, v := range weights {
		sum += v
		if v < min {
			min = v
		}
	}

	l, r := min, sum
	for l <= r {
		mid := (r-l)/2 + l
		if isOK(weights, mid, D) {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return l
}

func isOK(weight []int, w, D int) bool {
	tmp := 0
	for _, v := range weight {
		if tmp+v > w {
			D--
			// 注意
			if v > w {
				return false
			}
			tmp = v
		} else {
			tmp += v
		}
	}
	//if tmp > w {
	D--
	//}
	return D >= 0
}
