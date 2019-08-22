/*
 * @lc app=leetcode id=628 lang=golang
 *
 * [628] Maximum Product of Three Numbers
 */
//  (36 ms) √ Your runtime beats 96.23 %
func maximumProduct(nums []int) int {
	// 遍历过程中，保存三个最小负数，三个最大正数
	min1, min2 := 1001, 1001
	max1, max2, max3 := -1001, -1001, -1001 //3个 正数 1 <= 2 <= 3
	for _, v := range nums {
		if v > max3 {
			max3, max2, max1 = v, max3, max2
		} else if v > max2 {
			max2, max1 = v, max2
		} else if v > max1 {
			max1 = v
		}
		if v < min1 {
			min1, min2 = v, min1
		} else if v < min2 {
			min2 = v
		}

	}
	if max1*max2*max3 > min1*min2*max3 {
		return max1 * max2 * max3
	}
	return min1 * min2 * max3
}
