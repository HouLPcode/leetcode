/*
 * @lc app=leetcode.cn id=42 lang=golang
 *
 * [42] 接雨水
 *
 * https://leetcode-cn.com/problems/trapping-rain-water/description/
 *
 * algorithms
 * Hard (46.21%)
 * Likes:    561
 * Dislikes: 0
 * Total Accepted:    25.4K
 * Total Submissions: 55K
 * Testcase Example:  '[0,1,0,2,1,0,1,3,2,1,2,1]'
 *
 * 给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。
 * 
 * 
 * 
 * 上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。 感谢
 * Marcos 贡献此图。
 * 
 * 示例:
 * 
 * 输入: [0,1,0,2,1,0,1,3,2,1,2,1]
 * 输出: 6
 * 
 */
 //  4 ms, faster than 88.04%
// 一列一列计算雨水
func trap(height []int) int {
	if len(height) == 0 {
		return 0
	}
    // [i]表示前i个元素中的最大值，包含第i个元素
    maxLeft := make([]int, len(height))
    maxLeft[0] = height[0]
    for i :=1; i<len(height);i++ {
        if height[i] > maxLeft[i-1] {
            maxLeft[i] = height[i]
        } else {
            maxLeft[i] = maxLeft[i-1]
        }
    }
    // [i] 右边i个元素的最大值，包含第i个元素
    maxRight := make([]int, len(height))
    maxRight[len(height)-1] = height[len(height)-1]
    for i:=len(height)-2; i>=0; i-- {
        if height[i] > maxRight[i+1] {
            maxRight[i] = height[i]
        } else {
            maxRight[i] = maxRight[i+1]
        }
    }
    
    sum := 0 
    for i:=1; i<len(height)-1; i++ {// 首位不可能有雨水
        if height[i] < maxLeft[i-1] && height[i] < maxRight[i+1]{
            sum += min(maxLeft[i-1], maxRight[i+1]) - height[i]
        }
    }
    return sum
}
func min(a, b int) int{
    if a < b {
        return a
    }
    return b
}

