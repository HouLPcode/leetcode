/*
 * @lc app=leetcode id=747 lang=golang
 *
 * [747] Largest Number At Least Twice of Others
 */
// (0 ms)
// [0,0,0,1]
// [1]
func dominantIndex(nums []int) int {
	// 找前两个最大数，全是正数
	max1, max2 := -1, -1
	index := -1
	for k, v := range nums {
		if v > max1 {
			max1, max2 = v, max1
			index = k
		} else if v > max2 {
			max2 = v
		}
	}
	if max2 == -1 {
		// 只有一个数
		return index
	}
	if max1 != -1 && max2 == 0 { // 注意max2可能是0
		return index
	}
	if max1/max2 >= 2 {
		return index
	}
	return -1
}
