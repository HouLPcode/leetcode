/*
 * @lc app=leetcode id=11 lang=golang
 *
 * [11] Container With Most Water
 */
func maxArea(height []int) int {
	// O(n^2) 爆搜，统计过程中的最大值
	// tag: 2指针,左右两侧往中间遍历,哪侧低移动哪侧
	maxArea, area := 0, 0
	l, r := 0, len(height)-1
	for l < r{
		if height[r] > height[l]{ // 左侧小，l++
			area = (r-l)*height[l]
			l++
		}else{
			area = (r-l)*height[r] // 右侧小，r--
			r--
		}
		if area > maxArea{
			maxArea = area
		}
	}
	return maxArea
}

